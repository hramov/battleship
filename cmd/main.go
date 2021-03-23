package main

import (
	"encoding/json"
	"fmt"

	b "github.com/hramov/battleship/pkg/battlefield"
	connection "github.com/hramov/battleship/pkg/connection"
	ship "github.com/hramov/battleship/pkg/ship"
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
				client.CreateField()
				client.DrawField()
			})
			s.On("placeShip", func(_ string) {
				sh := ship.Ship{}
				sh.CreateShip()
				data, err := json.Marshal(sh)
				if err != nil {
					fmt.Println(err)
				}
				s.Emit("sendShip", string(data))
			})

			s.On("wrongShip", func(data string) {
				fmt.Println(data)
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
