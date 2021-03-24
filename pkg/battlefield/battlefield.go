package client

import (
	"fmt"
)

const FIELD_WIDTH = 12
const FIELD_HEIGHT = 12
const LETTER_STRING = "   А Б В Г Д Е Ж З И К\t\t   А Б В Г Д Е Ж З И К\n"

type Field [FIELD_WIDTH][FIELD_HEIGHT]string

type BattleField struct {
	Field     Field
	ShotField Field
}

func (b *BattleField) CreateField() {

	for i := 0; i < FIELD_HEIGHT; i++ {
		for j := 0; j < FIELD_WIDTH; j++ {
			if i == 0 || i == FIELD_HEIGHT-1 {
				b.Field[i][j] = "*"
				b.ShotField[i][j] = "*"
				continue
			}
			if j == 0 || j == FIELD_WIDTH-1 {
				b.Field[i][j] = "*"
				b.ShotField[i][j] = "*"
			} else {
				b.Field[i][j] = "_"
				b.ShotField[i][j] = "_"
			}
		}
	}
}

func (b *BattleField) DrawField() {
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
				fmt.Printf("|%s", b.Field[i][j])
			} else {
				fmt.Printf("|%s|", b.Field[i][j])
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
				fmt.Printf("|%s", b.ShotField[i][j])
			} else {
				fmt.Printf("|%s|", b.ShotField[i][j])
			}
		}
		fmt.Println()
	}
}
