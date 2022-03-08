package controller

import (
	"elton-okawa/battleship/internal/database"
	"elton-okawa/battleship/internal/use_case"
	"fmt"
)

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
