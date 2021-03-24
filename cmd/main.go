package main

import (
	"encoding/json"
	"fmt"

	b "github.com/hramov/battleship/pkg/battlefield"
	connection "github.com/hramov/battleship/pkg/connection"
	"github.com/hramov/battleship/pkg/ship"
	"github.com/hramov/battleship/pkg/utils"
)

func main() {

	client := b.Client{}
	s := connection.Execute("tcp", "127.0.0.1", "5000")

	handlers := make(map[string]func(data string))

	handlers["connect"] = func(data string) {
		fmt.Println("Connected!")
	}

	handlers["whoami"] = func(data string) {
		client.ID = data
		s.Emit("sendName", "BattleShip")
	}

	handlers["enemy"] = func(data string) {
		client.EnemyID = data
	}

	handlers["drawField"] = func(data string) {
		client.CreateField()
		client.DrawField()
	}

	handlers["placeShip"] = func(_ string) {
		utils.Log("PLACE SHIP")
		sh := ship.Ship{}
		sh.CreateShip()
		data, err := json.Marshal(sh)
		if err != nil {
			fmt.Println(err)
		}
		s.Emit("sendShip", string(data))
	}

	s.On(&handlers)

}
