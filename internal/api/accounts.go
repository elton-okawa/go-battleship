package api

import (
	"elton-okawa/battleship/internal/interface_adapter/presenter/rest"

	"github.com/labstack/echo/v4"
)

// Mapping from echo context -> actual arguments is the responsability of the controller
func (b *BattleshipImpl) CreateAccount(ctx echo.Context) error {
	rp := rest.New()

	// var postBody CreateAccountJSONBody
	// err := ctx.Bind(&postBody)
	// if err != nil {
	// 	restPresenter.CreateError(
	// 		http.StatusBadRequest,
	// 		"Invalid format for createAccount",
	// 	)
	// 	return restPresenter.Error()
	// }

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

	// var body AccountLoginJSONBody
	// if err := ctx.Bind(&body); err != nil {
	// 	restPresenter.CreateError(http.StatusBadRequest, "Invalid body to perform login")
	// 	return restPresenter.Error()
	// }

	// b.accounts.Login(restPresenter, body.Login, body.Password)

	// if restPresenter.Error() != nil {
	// 	return restPresenter.Error()
	// }

	// return ctx.JSON(restPresenter.Code(), restPresenter.Body())
}
