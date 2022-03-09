package controller

import (
	"elton-okawa/battleship/internal/database"
	"elton-okawa/battleship/internal/use_case"
	"fmt"
)

// TODO who should instantiate it?
var game use_case.Game = use_case.Game{
	Persistence: &database.DefaultDatabase,
}

func GetGame(id string) {
	fmt.Println("Get games")
}

func PostGame() *use_case.GameState {
	// TODO handle error
	gs, _ := game.Start()
	return gs
}

func Shoot(id string, row int, col int) (bool, int, *use_case.GameState) {
	// TODO handle error
	hit, ships, gameState, _ := game.Shoot(id, row, col)

	return hit, ships, gameState
}
