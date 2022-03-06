package api

import (
	"elton-okawa/battleship/internal/use_case"
	"fmt"
)

func GetGame(id string) {
	fmt.Println("Get games")
}

func PostGame() use_case.GameState {
	return use_case.StartGame()
}
