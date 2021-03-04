package ship

import (
	"fmt"
)

type Ship struct {
	Player     string
	Length     int
	StartX     int
	StartY     int
	Direction  int
	Hit        bool
	LivePoints int
	Live       bool
}

type Shot struct {
	X int
	Y int
}

func (s Ship) CreateShip(PlayerID string) Ship {
	s.Player = PlayerID
	fmt.Printf("%s ", "Введите числовую позицию (1-10):")
	fmt.Scanf("%d", &s.StartY)
	fmt.Printf("%s ", "Введите буквенную позицию (1-10):")
	fmt.Scanf("%d", &s.StartX)
	fmt.Printf("%s ", "Введите направление (0: -, 1: |):")
	fmt.Scanf("%d", &s.Direction)
	fmt.Printf("%s ", "Введите длину корабля (1,2,3,4):")
	fmt.Scanf("%d", &s.Length)
	s.LivePoints = s.Length
	fmt.Println(s)
	return s
}

func HitShip() Shot {
	var shot Shot
	fmt.Println("Введите координаты выстрела! Число:")
	fmt.Scanf("%d", &shot.X)
	fmt.Println("Введите координаты выстрела! Буква:")
	fmt.Scanf("%d", &shot.Y)
	return shot
}
