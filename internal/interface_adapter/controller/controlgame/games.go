package controlgame

import (
	"elton-okawa/battleship/internal/entity/jwttoken"
	"elton-okawa/battleship/internal/interface_adapter/controller"
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

func (gc Controller) PostGame(p ucgame.GameOutputBoundary, ctx controller.Context) error {
	claim, ok := ctx.Get("user").(jwttoken.Claim)
	if !ok {
		fmt.Printf("%+v\n", ctx.Get("user"))
		fmt.Println("error")
		// TODO return error of invalid claim
	}

	return gc.useCase.Start(p, claim.Player)
}

func (gc Controller) Shoot(p ucgame.GameOutputBoundary, id string, row int, col int) {
	gc.useCase.Shoot(p, id, row, col)
}
