package Api

import (
	"github.com/labstack/echo/v4"
)

type BattleshipImpl struct{}

func (*BattleshipImpl) CreateAccount(ctx echo.Context) error {

	return nil
}

func SetupHandler() *echo.Echo {
	var app BattleshipImpl
	e := echo.New()
	RegisterHandlers(e, &app)

	return e
}
