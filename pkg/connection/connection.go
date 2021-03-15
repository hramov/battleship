package connection

import (
	"bufio"
	"log"
	"net"
)

type Socket struct {
	protocol  string
	host_ip   string
	host_port string
	conn      net.Conn
}

func Execute(protocol, ip, port string) {
	socket := Socket{protocol, ip, port, nil}
	conn := socket.ConnectToServer()
	socket.conn = conn
	for {
		socket.MaintainConnection()
	}
}

func (s *Socket) ConnectToServer() net.Conn {
	conn, err := net.Dial(s.protocol, s.host_ip+":"+s.host_port)
	if err != nil {
		log.Fatal("Ошибка при подключении к серверу!")
	}
	return conn
}

func (s *Socket) On(rawEvent string, callback func(data string)) {
	rawData, _ := bufio.NewReader(s.conn).ReadString('\n')
	event, data := split(rawData, ":")

	if event == rawEvent {
		callback(data)
	}
}

func (s *Socket) Emit(event string, data string) {
	s.conn.Write([]byte(event + ":" + data + "\n"))
}

func (s *Socket) MaintainConnection() {
	s.On("whoami", func(data string) {
		s.Emit("sendName", "Battleship")
	})
}

func split(message string, delim string) (string, string) {
	var breakPosition int = 0
	var charArray []string
	var eventString, dataString string
	for i, char := range message {
		charArray = append(charArray, string(char))
		if string(char) == delim {
			breakPosition = i
		}
	}
	event := charArray[:breakPosition]
	data := charArray[breakPosition+1:]
	for i := 0; i < len(event); i++ {
		eventString += string(event[i])
	}
	for i := 0; i < len(data); i++ {
		dataString += string(data[i])
	}
	return eventString, dataString
}
