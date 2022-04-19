package router

import (
	"elton-okawa/battleship/internal/interface_adapter/presenter/rest"

	"github.com/labstack/echo/v4"
)

func (b *BattleshipImpl) CreateAccount(ctx echo.Context) error {
	rp := rest.New()

	if err := b.accounts.CreateAccount(rp, ctx); err != nil {
		return err
	}

	return ctx.JSON(rp.Code(), rp.Body())
}

func (b *BattleshipImpl) AccountLogin(ctx echo.Context) error {
	rp := rest.New()

	if err := b.accounts.Login(rp, ctx); err != nil {
		return err
	}

	return ctx.JSON(rp.Code(), rp.Body())
}
