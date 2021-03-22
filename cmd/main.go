package main

import (
	"fmt"

	b "github.com/hramov/battleship/pkg/battlefield"
	connection "github.com/hramov/battleship/pkg/connection"
)

func main() {

	client := b.Client{}

	connection.Execute("tcp", "127.0.0.1", "5000", func(s *connection.Socket) {
		for {
			s.On("connect", func(data string) {
				fmt.Println("Connected!")
			})
			s.On("whoami", func(data string) {
				client.ID = data
				s.Emit("sendName", "BattleShip")
			})
			s.On("enemy", func(data string) {
				client.EnemyID = data
			})
			s.On("drawField", func(data string) {
				client.DrawField(data)
			})
			s.On("placeShip", func(data string) {
				//
			})
			s.On("fire", func(data string) {
				//
			})
			s.On("hit", func(data string) {
				//
			})
			s.On("dead", func(data string) {
				//
			})
			s.On("end", func(data string) {
				//
			})
		}
	})

}
