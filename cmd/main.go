package main

// import (
// 	// "github.com/hramov/battleship/pkg/menu"
// 	"github.com/zhouhui8915/go-socket.io-client"
// )

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"

	bf "github.com/hramov/battleship/pkg/battlefield"
	"github.com/hramov/battleship/pkg/ship"
	"github.com/hramov/battleship/pkg/socket"
)

func main() {

	client := socket.CreateSocket()
	var ClientID string

	client.On("connection", func() {
		log.Printf("Успешно подключился к серверу\n")
	})

	client.On("whoami", func(ID string) {
		ClientID = ID
	})

	client.On("updateField", func(data []byte) {
		var client bf.Client
		json.Unmarshal(data, &client)
		bf.DrawField(client)
	})

	client.On("placeShips", func() {
		var ships []ship.Ship
		i := 0
		s := ship.Ship{}

		for i < 10 {
			ships = append(ships, s.CreateShip(ClientID))
			i++
		}

		data, _ := json.Marshal(ships)
		client.Emit("sendShips", data)
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
		log.Printf("Отключился от сервера\n")
	})

	//Make this work without stop
	reader := bufio.NewReader(os.Stdin)
	for {
		data, _, _ := reader.ReadLine()
		fmt.Println(data)
	}
}
