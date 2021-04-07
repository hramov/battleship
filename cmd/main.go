package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	battlefield "github.com/hramov/battleship/pkg/battlefield"
	connection "github.com/hramov/battleship/pkg/connection"
	"github.com/hramov/battleship/pkg/ship"
	"github.com/hramov/battleship/pkg/shot"
	"github.com/hramov/battleship/pkg/utils"
)

func main() {

	b := battlefield.BattleField{}
	c := connection.Client{}
	s := connection.Execute("tcp", "127.0.0.1", "5000")

	handlers := make(map[string]func(data string, client *connection.Client))

	handlers["connect"] = func(data string, client *connection.Client) {
		fmt.Println("Connected!")
	}

	handlers["whoami"] = func(data string, client *connection.Client) {
		newID, _ := strconv.Atoi(strings.TrimSuffix(data, "\n"))
		(*client).ID = newID
	}

	handlers["enemy"] = func(data string, client *connection.Client) {
		(*client).EnemyID, _ = strconv.Atoi(strings.TrimSuffix(data, "\n"))
	}

	handlers["drawField"] = func(data string, client *connection.Client) {
		b.CreateField()
		json.Unmarshal([]byte(data), &b)
		b.DrawField()
	}

	handlers["updateField"] = func(data string, client *connection.Client) {
		utils.Log("Update")
		b.UpdateField(data)
	}

	handlers["placeShip"] = func(_ string, client *connection.Client) {
		sh := ship.Ship{}
		sh.CreateShip()
		data, err := json.Marshal(sh)
		if err != nil {
			fmt.Println(err)
		}
		s.Emit("sendShip", string(data))
	}

	handlers["wrongShip"] = func(data string, client *connection.Client) {
		utils.Log(data)
	}

	handlers["rightShip"] = func(data string, client *connection.Client) {
		utils.Log(data)
	}

	handlers["makeShot"] = func(data string, client *connection.Client) {
		turn, err := strconv.ParseBool(strings.TrimSuffix(data, "\n"))
		if err != nil {
			fmt.Println(err)
		}
		if turn {
			newShot := shot.Shot{}
			newShot.MakeShot()
			shotData, err := json.Marshal(newShot)
			if err != nil {
				utils.Log(err.Error())
			}
			s.Emit("shot", string(shotData))
		} else {
			fmt.Println("Ожидайте хода противника...")
		}
	}

	handlers["wrongShot"] = func(data string, client *connection.Client) {
		utils.Log(data)
	}

	handlers["hit"] = func(data string, client *connection.Client) {
		b.DrawShot(data, "X")
	}

	handlers["missed"] = func(data string, client *connection.Client) {
		b.DrawShot(data, "*")
	}

	s.On(&handlers, &c)
}
