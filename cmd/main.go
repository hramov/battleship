package main

import (
	"fmt"

	connection "github.com/hramov/battleship/pkg/connection"
)

func main() {

	connection.Execute("tcp", "127.0.0.1", "5000", func(s *connection.Socket) {

		s.On("connect", func(data string) {
			fmt.Println("Connected!")
		})

		s.On("whoami", func(data string) {
			fmt.Println("Who am I?")
			s.Emit("sendName", "BattleShip")
		})

	})

}
