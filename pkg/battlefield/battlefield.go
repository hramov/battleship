package battlefield

import (
	"fmt"

	"github.com/hramov/battleship/pkg/ship"
	"github.com/hramov/battleship/pkg/utils"
)

const FIELD_WIDTH = 10
const FIELD_HEIGHT = 10

type Field [FIELD_WIDTH][FIELD_HEIGHT]string

type BattleField struct {
	myField    Field
	enemyField Field
}

func (b BattleField) CreateField() BattleField {
	for i := 0; i < FIELD_WIDTH; i++ {
		for j := 0; j < FIELD_HEIGHT; j++ {
			b.myField[i][j] = "_"
		}
	}
	for i := 0; i < FIELD_WIDTH; i++ {
		for j := 0; j < FIELD_HEIGHT; j++ {
			b.enemyField[i][j] = "_"
		}
	}
	return b
}

func (b BattleField) DrawField() {
	fmt.Printf("   А Б В Г Д Е Ж З И К\t\t   А Б В Г Д Е Ж З И К\n")
	for i := 0; i < FIELD_WIDTH; i++ {
		if i != FIELD_WIDTH-1 {
			fmt.Printf(" %d", i+1)
		} else {
			fmt.Printf("%d", i+1)
		}
		for j := 0; j < FIELD_HEIGHT; j++ {
			if j != FIELD_HEIGHT-1 {
				fmt.Printf("|%s", b.myField[i][j])
			} else {
				fmt.Printf("|%s|", b.myField[i][j])
			}
		}
		fmt.Printf("\t\t")
		if i != FIELD_WIDTH-1 {
			fmt.Printf(" %d", i+1)
		} else {
			fmt.Printf("%d", i+1)
		}
		for j := 0; j < FIELD_HEIGHT; j++ {
			if j != FIELD_HEIGHT-1 {
				fmt.Printf("|%s", b.enemyField[i][j])
			} else {
				fmt.Printf("|%s|", b.enemyField[i][j])
			}
		}
		fmt.Println()
	}
}

func (b BattleField) UpdateField(s ship.Ship) BattleField {
	fmt.Printf("\nХод: %s-%d\n", utils.Parser(s.StartX), s.StartY)
	for i := 0; i < s.Length; i++ {
		if s.Direction == 0 {
			b.enemyField[(s.StartX - 1)][(s.StartY-1)+i] = "X"
		} else if s.Direction == 1 {
			b.enemyField[(s.StartX-1)+i][(s.StartY - 1)] = "X"
		}
	}
	b.DrawField()
	return b
}

func (b BattleField) ClearField() {
	b.myField = [10][10]string{}
	b.enemyField = [10][10]string{}
	b.DrawField()
}

func (b BattleField) CheckLength(StartX, StartY, Direction int) int {

	return 0
}
