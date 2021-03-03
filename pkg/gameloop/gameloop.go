package gameloop

import (
	"fmt"

	"github.com/hramov/battleship/pkg/battlefield"
	"github.com/hramov/battleship/pkg/ship"
)

func Start(Player bool) {
	//Creating and drawing battlefield
	b := battlefield.BattleField{}
	b = b.CreateField()
	b.DrawField()

	//Creating and drawing ships
	var ships []ship.Ship

	PlaceShips(Player, &ships, &b)
	PlaceShips(!Player, &ships, &b)

	Game(Player, &b, &ships)
}

func PlaceShips(Player bool, ships *[]ship.Ship, b *battlefield.BattleField) {
	i := 0
	for i < 5 {
		fmt.Printf("Корабль %d:\n", i+1)
		s := ship.Ship{}
		s = s.CreateShip(Player)
		_, err := b.CheckShip(s)
		if err != nil {
			fmt.Println(err)
		} else {
			*ships = append(*ships, s)
			fmt.Println(*ships)
			*b = b.DrawShip(s)
			i++
		}
	}
}

func Game(Player bool, b *battlefield.BattleField, ships *[]ship.Ship) {
	for {
		ShotX, ShotY := ship.HitShip()

		err := b.CheckShot(ShotX, ShotY)
		if err != nil {
			fmt.Println(err)
			continue
		}

		err1 := b.CheckHit(Player, ShotX, ShotY, ships)
		if err1 != nil {
			b.DrawShot(ShotX, ShotY, 0)
		} else {
			b.DrawShot(ShotX, ShotY, 1)
			continue
		}

		Player = !Player

	}
}
