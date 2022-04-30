package controlgame

import (
	"elton-okawa/battleship/internal/usecase/ucgame"
	"fmt"
)

func New(g ucgame.UseCase) Controller {
	return Controller{
		useCase: g,
	}
}

type Controller struct {
	useCase ucgame.UseCase
}

func (gc Controller) GetGame(id string) {
	fmt.Println("Get games")
}

func (gc Controller) PostGame(p ucgame.GameOutputBoundary) {
	gc.useCase.Start(p, "")
}

func (gc Controller) Shoot(p ucgame.GameOutputBoundary, id string, row int, col int) {
	gc.useCase.Shoot(p, id, row, col)
}
