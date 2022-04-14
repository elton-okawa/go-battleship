package Api

import (
	"elton-okawa/battleship/internal/interface_adapter/presenter/rest"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (b *BattleshipImpl) CreateAccount(ctx echo.Context) error {
	restPresenter := rest.NewRestApiPresenter(ctx)

	var postBody CreateAccountJSONBody
	err := ctx.Bind(&postBody)
	if err != nil {
		restPresenter.SendError(
			http.StatusBadRequest,
			"Invalid format for createAccount",
		)
		return restPresenter.Error()
	}

	b.accounts.CreateAccount(restPresenter, postBody.Login, postBody.Password)
	return restPresenter.Error()
}

func (b *BattleshipImpl) AccountLogin(ctx echo.Context) error {
	restPresenter := rest.NewRestApiPresenter(ctx)

	var postBody AccountLoginJSONBody
	if err := ctx.Bind(&postBody); err != nil {
		restPresenter.SendError(http.StatusBadRequest, "Invalid body to perform login")
		return restPresenter.Error()
	}

	b.accounts.Login(restPresenter, postBody.Login, postBody.Password)
	return restPresenter.Error()
}
