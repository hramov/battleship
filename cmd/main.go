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

	bclient "github.com/hramov/battleship/pkg/client"
	"github.com/hramov/battleship/pkg/socket"
)

func main() {

	client := socket.CreateSocket()

	client.On("connection", func() {
		log.Printf("Успешно подключился к серверу\n")
	})

	client.On("field", func(data []byte) {

		var client bclient.Client
		err := json.Unmarshal(data, &client)

		if err != nil {
			fmt.Println("Ошибка при декодировании JSON")
		}

		bclient.DrawField(client)
	})

	client.On("shot", func() {
		shot := bclient.MakeShot()
		client.Emit(shot)
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

	reader := bufio.NewReader(os.Stdin)
	for {
		data, _, _ := reader.ReadLine()
		shot := string(data)
		client.Emit("shot", shot)
		log.Printf("Ваш ход: %v\n", shot)
	}
}
