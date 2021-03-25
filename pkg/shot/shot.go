package shot

import "fmt"

type Shot struct {
	X int
	Y int
}

func (s *Shot) MakeShot() {
	fmt.Println("Введите числовую координату выстрела: ")
	fmt.Scanf("%d", &s.X)
	fmt.Println("Введите буквенную координату выстрела: ")
	fmt.Scanf("%d", &s.Y)
}
