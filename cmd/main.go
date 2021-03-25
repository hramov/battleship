package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	battlefield "github.com/hramov/battleship/pkg/battlefield"
	connection "github.com/hramov/battleship/pkg/connection"
	"github.com/hramov/battleship/pkg/ship"
	"github.com/hramov/battleship/pkg/shot"
	"github.com/hramov/battleship/pkg/utils"
)

func main() {

	c := connection.Client{}
	b := battlefield.BattleField{}

	s := connection.Execute("tcp", "127.0.0.1", "5000")

	handlers := make(map[string]func(data string))

	handlers["connect"] = func(data string) {
		fmt.Println("Connected!")
	}

	handlers["whoami"] = func(data string) {
		c.ID, _ = strconv.Atoi(data)
		s.Emit("sendName", "BattleShip")
	}

	handlers["enemy"] = func(data string) {
		c.EnemyID, _ = strconv.Atoi(data)
	}

	handlers["drawField"] = func(data string) {
		b.CreateField()
		json.Unmarshal([]byte(data), &b)
		b.DrawField()
	}

	handlers["placeShip"] = func(_ string) {
		sh := ship.Ship{}
		sh.CreateShip()
		data, err := json.Marshal(sh)
		if err != nil {
			fmt.Println(err)
		}
		s.Emit("sendShip", string(data))
	}

	handlers["wrongShip"] = func(data string) {
		utils.Log(data)
	}

	handlers["rightShip"] = func(data string) {
		utils.Log(data)
	}

	handlers["makeShot"] = func(data string) {
		if data == strconv.Itoa(c.ID) {
			newShot := shot.Shot{}
			newShot.MakeShot()
			shotData, err := json.Marshal(newShot)
			if err != nil {
				utils.Log(err.Error())
			}
			s.Emit("shot", string(shotData))
		}
	}

	handlers["wrongShot"] = func(data string) {
		utils.Log(data)
	}

	handlers["rightShot"] = func(data string) {
		b.DrawShot(data)
	}

	s.On(&handlers)

}
