package main

import (
	"encoding/json"
	"fmt"
	"log"

	b "github.com/hramov/battleship/pkg/battlefield"
	ship "github.com/hramov/battleship/pkg/ship"
	socket "github.com/hramov/battleship/pkg/socket"
)

func main() {
	var ClientID string
	var isGame bool = true
	var player b.Client
	client := socket.CreateSocket()

	client.On("connection", func() {
		log.Printf("Успешно подключился к серверу\n")
	})

	client.On("whoami", func(ID string) {
		ClientID = ID
	})

	client.On("updateField", func(data []byte) {
		fmt.Println("123")
		json.Unmarshal(data, &player)
		b.DrawField(player)
	})

	client.On("placeShip", func() {
		s := ship.Ship{}
		s = s.CreateShip(ClientID)
		data, _ := json.Marshal(s)
		client.Emit("sendShip", data)
	})

	client.On("makeShot", func() {
		shot := ship.HitShip()
		data, _ := json.Marshal(shot)
		client.Emit("sendShot", data)
	})

	client.On("message", func(message string) {
		fmt.Println(message)
	})

	client.On("error", func() {
		log.Printf("Ошибка подключения к серверу\n")
	})

	client.On("disconnection", func() {
		isGame = false
		log.Printf("Отключился от сервера\n")
	})

	for isGame == true {
	}
}
