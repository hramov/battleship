package gameloop

import (
	"fmt"

	"github.com/hramov/battleship/pkg/battlefield"
	"github.com/hramov/battleship/pkg/ship"
)

func PlaceShips() {

	//Creating and drawing battlefield
	b := battlefield.BattleField{}
	b = b.CreateField()
	b.DrawField()

	//Creating and drawing ships
	var ships []ship.Ship

	for i := 0; i < 10; i++ {
		fmt.Printf("Корабль %d:\n", i+1)
		s := ship.Ship{}
		s = s.CreateShip()

		_, err := b.CheckShip(s)

		if err != nil {
			fmt.Println(err)
		} else {
			ships = append(ships, s)
			b = b.UpdateField(s)
			fmt.Println(b)
		}
	}

}
