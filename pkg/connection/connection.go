package connection

import (
	"bufio"
	"log"
	"net"
	"time"

	"github.com/hramov/battleship/pkg/utils"
)

const (
	socketDelay = 100
)

type Client struct {
	ID      int
	EnemyID int
}

type Socket struct {
	protocol  string
	host_ip   string
	host_port string
	conn      net.Conn
	from      chan string
	to        chan string
}

func Execute(protocol, ip, port string) *Socket {
	socket := Socket{protocol, ip, port, nil, make(chan string), make(chan string)}
	conn := socket.ConnectToServer()
	socket.conn = conn
	socket.maintainConnections()
	return &socket
}

func (s *Socket) ConnectToServer() net.Conn {
	conn, err := net.Dial(s.protocol, s.host_ip+":"+s.host_port)
	if err != nil {
		log.Fatal("Ошибка при подключении к серверу!")
	}
	return conn
}

func (s *Socket) listen() {
	for {
		rawData, err := bufio.NewReader(s.conn).ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		s.from <- rawData
	}
}

func (s *Socket) On(handlers *map[string]func(data string)) {

	for {
		time.Sleep(time.Second / socketDelay)
		rawData := <-s.from
		rawEvent, data := utils.Split(rawData, "|")
		for event, handler := range *handlers {
			if event == rawEvent {
				handler(data)
			}
		}
	}
}

func (s *Socket) speak() {
	for {
		event, data := utils.Split(<-s.to, "|")
		s.conn.Write([]byte(string(event) + "|" + string(data) + "\n"))
	}
}

func (s *Socket) Emit(event string, data string) {
	time.Sleep(time.Second / socketDelay)
	s.to <- event + "|" + data
}

func (s *Socket) maintainConnections() {
	go s.speak()
	go s.listen()
}
