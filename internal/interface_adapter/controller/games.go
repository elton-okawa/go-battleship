package controller

import (
	"elton-okawa/battleship/internal/use_case"
	"fmt"
)

func NewGamesController(g *use_case.Game) *GamesController {
	return &GamesController{
		game: g,
	}
}

type GamesController struct {
	game *use_case.Game
}

func (gc *GamesController) GetGame(id string) {
	fmt.Println("Get games")
}

func (gc *GamesController) PostGame() {
	gc.game.Start()
}

func (gc *GamesController) Shoot(id string, row int, col int) {
	gc.game.Shoot(id, row, col)
}
