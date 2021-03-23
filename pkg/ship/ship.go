package ship

import (
	"fmt"
)

type Ship struct {
	Length     int  `json: "length"`
	StartX     int  `json: "startX"`
	StartY     int  `json: "startY"`
	Direction  int  `json: "direction"`
	Hit        bool `json: "hit"`
	LivePoints int  `json: "livePoints"`
	Live       bool `json: "live"`
}

type Shot struct {
	X int
	Y int
}

func (s *Ship) CreateShip() {
	fmt.Printf("%s ", "Введите числовую позицию (1-10):")
	fmt.Scanf("%d", &s.StartY)
	fmt.Printf("%s ", "Введите буквенную позицию (1-10):")
	fmt.Scanf("%d", &s.StartX)
	fmt.Printf("%s ", "Введите направление (0: -, 1: |):")
	fmt.Scanf("%d", &s.Direction)
	fmt.Printf("%s ", "Введите длину корабля (1,2,3,4):")
	fmt.Scanf("%d", &s.Length)
	s.LivePoints = s.Length
	s.Hit = false
	s.Live = true
}

func HitShip() Shot {
	var shot Shot
	fmt.Println("Введите координаты выстрела! Число:")
	fmt.Scanf("%d", &shot.X)
	fmt.Println("Введите координаты выстрела! Буква:")
	fmt.Scanf("%d", &shot.Y)
	return shot
}
