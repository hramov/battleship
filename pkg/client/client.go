package client

import (
	"fmt"
)

const FIELD_WIDTH = 12
const FIELD_HEIGHT = 12
const LETTER_STRING = "   А Б В Г Д Е Ж З И К\t\t   А Б В Г Д Е Ж З И К\n"

type Field [FIELD_WIDTH][FIELD_HEIGHT]string

type BattleField struct {
	myField    Field
	enemyField Field
}

type Client struct {
	ID    string
	Field BattleField
}

func DrawField(client Client) {

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
				fmt.Printf("|%s", client.Field.myField[i][j])
			} else {
				fmt.Printf("|%s|", client.Field.myField[i][j])
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
				fmt.Printf("|%s", client.Field.enemyField[i][j])
			} else {
				fmt.Printf("|%s|", client.Field.enemyField[i][j])
			}
		}
		fmt.Println()
	}
}

func MakeShot() string {
	fmt.Println("Введите координаты выстрела: (1А)")
	var shot string
	fmt.Scanf("%s", &shot)
	return shot
}
