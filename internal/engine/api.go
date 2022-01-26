package engine

import (
	"fmt"
)

func StartGame() {
	fmt.Println("Start game")

	board := Init()
	fmt.Println(board)
}
