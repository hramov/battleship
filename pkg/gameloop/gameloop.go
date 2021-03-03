package gameloop

import (
	"fmt"

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
	Player := true

	PlaceShips(Player, &ships, &b)
	PlaceShips(!Player, &ships, &b)

	Game(Player, &b, &ships)
}

func PlaceShips(Player bool, ships *[]ship.Ship, b *battlefield.BattleField) {
	i := 0
	fmt.Printf("Расстановка кораблей для игрока %t\n", Player)
	for i < 1 {
		fmt.Printf("Корабль %d:\n", i+1)
		s := ship.Ship{}
		s = s.CreateShip(Player)
		_, err := b.CheckShip(s)
		if err != nil {
			fmt.Println(err)
		} else {
			*ships = append(*ships, s)
			*b = b.DrawShip(s)
			i++
		}
	}
}

func Game(Player bool, b *battlefield.BattleField, ships *[]ship.Ship) {

	for {

		fmt.Printf("Играем с %t\n", Player)
		var ShotX, ShotY int
		fmt.Println("Введите координаты выстрела! Число:")
		fmt.Scanf("%d", &ShotX)
		fmt.Println("Введите координаты выстрела! Буква:")
		fmt.Scanf("%d", &ShotY)

		err := b.CheckShot(Player, ShotX, ShotY)
		if err != nil {
			fmt.Println(err)
			continue
		}

		err1 := b.CheckHit(Player, ShotX, ShotY, ships)
		if err1 != nil {
			fmt.Println(err1)
			*b = b.DrawShot(Player, ShotX, ShotY, 0)
			Player = !Player
		} else {
			fmt.Println("Попал!")
			*b = b.DrawShot(Player, ShotX, ShotY, 1)
		}

	}
}
