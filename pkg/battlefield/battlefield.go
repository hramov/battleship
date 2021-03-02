package battlefield

import (
	"fmt"

	"github.com/hramov/battleship/pkg/ship"
	"github.com/hramov/battleship/pkg/utils"
)

const FIELD_WIDTH = 12
const FIELD_HEIGHT = 12
const LETTER_STRING = "   А Б В Г Д Е Ж З И К\t\t   А Б В Г Д Е Ж З И К\n"

type Field [FIELD_WIDTH][FIELD_HEIGHT]string

type BattleField struct {
	myField    Field
	enemyField Field
}

func (b BattleField) CreateField() BattleField {
	for i := 0; i < FIELD_HEIGHT; i++ {
		for j := 0; j < FIELD_WIDTH; j++ {
			if i == 0 || i == FIELD_HEIGHT-1 {
				b.myField[i][j] = "*"
				b.enemyField[i][j] = "*"
				continue
			}
			if j == 0 || j == FIELD_WIDTH-1 {
				b.myField[i][j] = "*"
				b.enemyField[i][j] = "*"
			} else {
				b.myField[i][j] = "_"
				b.enemyField[i][j] = "_"
			}
		}
	}
	return b
}

func (b BattleField) DrawField() {

	fmt.Printf(LETTER_STRING)
	for i := 1; i < FIELD_HEIGHT-1; i++ {

		//My field drawing
		if i != FIELD_HEIGHT-2 {
			fmt.Printf(" %d", i)
		} else {
			fmt.Printf("%d", i)
		}
		for j := 1; j < FIELD_WIDTH-1; j++ {
			if j != FIELD_WIDTH-2 {
				fmt.Printf("|%s", b.myField[i][j])
			} else {
				fmt.Printf("|%s|", b.myField[i][j])
			}
		}
		fmt.Printf("\t\t")

		//Enemy field drawing
		if i != FIELD_WIDTH-2 {
			fmt.Printf(" %d", i)
		} else {
			fmt.Printf("%d", i)
		}
		for j := 1; j < FIELD_HEIGHT-1; j++ {
			if j != FIELD_HEIGHT-2 {
				fmt.Printf("|%s", b.enemyField[i][j])
			} else {
				fmt.Printf("|%s|", b.enemyField[i][j])
			}
		}
		fmt.Println()
	}
}

func (b BattleField) UpdateField(s ship.Ship) BattleField {
	fmt.Printf("\nХод: %s-%d\n", utils.Parser(s.StartY), s.StartX)
	for i := 0; i < s.Length; i++ {
		if s.Direction == 0 {
			b.enemyField[s.StartY][s.StartX+i] = "X"
		} else if s.Direction == 1 {
			b.enemyField[s.StartY+i][s.StartX] = "X"
		}
	}
	b.DrawField()
	return b
}

func (b BattleField) ClearField() {
	b.myField = Field{}
	b.enemyField = Field{}
	b.DrawField()
}

func (b BattleField) CheckShip(s ship.Ship) (bool, error) {

	errorMessage := "Начальное сообщение"

	if b.enemyField[s.StartY][s.StartX] == "_" {
		if s.Direction == 0 {
			if s.StartY+s.Length < FIELD_HEIGHT {
				if b.enemyField[s.StartY+s.Length][s.StartX] != "_" {
					errorMessage = "Уперся в *"
				} else {
					return true, nil
				}
			} else {
				errorMessage = "Вышел за границу"
			}
		} else {
			if s.StartX+s.Length < FIELD_WIDTH {
				if b.enemyField[s.StartY][s.StartX+s.Length] != "_" {
					errorMessage = "Уперся в *"
				} else {
					return true, nil
				}
			} else {
				errorMessage = "Вышел за границу"
			}
		}
	} else {
		errorMessage = "Первое условие"
	}
	return false, fmt.Errorf("%s", errorMessage)
}
