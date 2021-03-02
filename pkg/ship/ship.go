package ship

import (
	"fmt"
)

type Ship struct {
	Length    int
	StartX    int
	StartY    int
	Direction int
}

func (s Ship) CreateShip() Ship {

	fmt.Printf("%s ", "Введите буквенную позицию (1-10):")
	fmt.Scanf("%d", &s.StartX)
	fmt.Printf("%s ", "Введите числовую позицию (1-10):")
	fmt.Scanf("%d", &s.StartY)
	fmt.Printf("%s ", "Введите направление (0: -, 1: |):")
	fmt.Scanf("%d", &s.Direction)

	//field.CheckLength

	fmt.Printf("%s ", "Введите длину корабля (1,2,3,4):")
	fmt.Scanf("%d", &s.Length)

	return s
}

func HitShip() {

}

func DestroyShip(id int) {
	var ships [][]int
	ships[id] = nil
}
