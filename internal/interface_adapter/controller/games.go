package controller

import (
	"elton-okawa/battleship/internal/usecase/game"
	"fmt"
)

func NewGamesController(g game.GameUseCase) GamesController {
	return GamesController{
		useCase: g,
	}
}

type GamesController struct {
	useCase game.GameUseCase
}

func (gc GamesController) GetGame(id string) {
	fmt.Println("Get games")
}

func (gc GamesController) PostGame(p game.GameOutputBoundary) {
	gc.useCase.Start(p)
}

func (gc GamesController) Shoot(p game.GameOutputBoundary, id string, row int, col int) {
	gc.useCase.Shoot(p, id, row, col)
}
