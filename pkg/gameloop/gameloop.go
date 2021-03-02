package gameloop

import (
	"github.com/hramov/battleship/pkg/battlefield"
	"github.com/hramov/battleship/pkg/ship"
)

func Start() {

	//Creating and drawing battlefield
	b := battlefield.BattleField{}
	b = b.CreateField()
	b.DrawField()

	//Creating and drawing ships
	var ships []ship.Ship

	s := ship.Ship{}
	s = s.CreateShip()
	ships = append(ships, s)

	b = b.UpdateField(s)

}

func Turn() {

}

func Hit() {

}

func Destroy() {

}

func IsWon() {

}

func Stop() {

}
