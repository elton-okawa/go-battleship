package router

import (
	"elton-okawa/battleship/internal/interface_adapter/presenter/rest"

	"github.com/labstack/echo/v4"
)

func (b *BattleshipImpl) CreateGame(ctx echo.Context) error {
	restPresenter := rest.New()

	b.games.PostGame(restPresenter, ctx)
	// return restPresenter.Error()
	return nil
}

func (b *BattleshipImpl) GameShoot(ctx echo.Context, id string) error {
	return nil
	// restPresenter := rest.New()

	// var body GameShootJSONBody
	// if err := ctx.Bind(&body); err != nil {
	// 	restPresenter.CreateError(http.StatusBadRequest, "Invalid shoot body")
	// 	return restPresenter.Error()
	// }

	// b.games.Shoot(restPresenter, id, body.Row, body.Col)
	// return restPresenter.Error()
}
