package engine

import (
	"fmt"
)

// TODO persist it somewhere
var board Board

func StartGame() {
	fmt.Println("Start game")

	board = Init()
	fmt.Println(board)
}

func Shoot(row, col int) (bool, int, *Board) {
	hit, ships := board.Shoot(row, col)
	return hit, ships, &board
}
