package Api

import (
	"elton-okawa/battleship/internal/interface_adapter/presenter/rest"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (b *BattleshipImpl) CreateGame(ctx echo.Context) error {
	restPresenter := rest.NewRestApiPresenter(ctx)

	b.games.PostGame(restPresenter)
	return restPresenter.Error()
}

func (b *BattleshipImpl) GameShoot(ctx echo.Context, id string) error {
	restPresenter := rest.NewRestApiPresenter(ctx)

	var body GameShootJSONBody
	if err := ctx.Bind(&body); err != nil {
		restPresenter.SendError(http.StatusBadRequest, "Invalid shoot body")
		return restPresenter.Error()
	}

	b.games.Shoot(restPresenter, id, body.Row, body.Col)
	return restPresenter.Error()
}
