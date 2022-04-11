package Api

import (
	"elton-okawa/battleship/internal/interface_adapter/presenter/rest"

	"github.com/labstack/echo/v4"
)

func (b *BattleshipImpl) CreateGame(ctx echo.Context) error {
	restPresenter := rest.NewRestApiPresenter(ctx)

	b.games.PostGame(restPresenter)
	return restPresenter.Error()
}
