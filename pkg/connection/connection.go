package connection

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"sync"

	"github.com/hramov/battleship/pkg/utils"
)

type Socket struct {
	protocol  string
	host_ip   string
	host_port string
	conn      net.Conn
	from      chan string
	to        chan string
}

func Execute(protocol, ip, port string, handler func(s *Socket)) {
	socket := Socket{protocol, ip, port, nil, make(chan string, 10), make(chan string, 10)}
	conn := socket.ConnectToServer()
	socket.conn = conn
	socket.maintainConnections(handler)
}

func (s *Socket) ConnectToServer() net.Conn {
	conn, err := net.Dial(s.protocol, s.host_ip+":"+s.host_port)
	if err != nil {
		log.Fatal("Ошибка при подключении к серверу!")
	}
	return conn
}

func (s *Socket) listen() {
	rawData, err := bufio.NewReader(s.conn).ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	s.from <- rawData
}

func (s *Socket) On(rawEvent string, callback func(data string)) {

	rawData := <-s.from
	fmt.Println(rawData)
	event, data := utils.Split(rawData, ":")
	if event == rawEvent {
		callback(string(data))
	}
}

func (s *Socket) speak() {
	event, data := utils.Split(<-s.to, ":")
	s.conn.Write([]byte(string(event) + ":" + string(data) + "\n"))
}

func (s *Socket) Emit(event string, data string) {
	s.to <- event + ":" + data
}

func (s *Socket) maintainConnections(handler func(s *Socket)) {
	wg := sync.WaitGroup{}
	wg.Add(4)

	go func() {
		go s.listen()
		go s.speak()
	}()

	go handler(s)

	wg.Wait()
}
