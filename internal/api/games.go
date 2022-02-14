package api

import (
	"elton-okawa/battleship/internal/engine"
	"fmt"
)

func GetGame(id string) {
	fmt.Println("Get games")
}

func PostGame() engine.GameState {
	return engine.StartGame()
}
