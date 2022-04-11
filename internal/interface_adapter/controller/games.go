package controller

import (
	"elton-okawa/battleship/internal/use_case/game"
	"fmt"
)

func NewGamesController(g *game.Game) *GamesController {
	return &GamesController{
		game: g,
	}
}

type GamesController struct {
	game *game.Game
}

func (gc *GamesController) GetGame(id string) {
	fmt.Println("Get games")
}

func (gc *GamesController) PostGame(p game.GameOutputBoundary) {
	gc.game.Start(p)
}

func (gc *GamesController) Shoot(p game.GameOutputBoundary, id string, row int, col int) {
	gc.game.Shoot(p, id, row, col)
}
