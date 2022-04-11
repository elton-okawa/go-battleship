package Api

import (
	"elton-okawa/battleship/internal/interface_adapter/presenter/rest"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (b *BattleshipImpl) CreateAccount(ctx echo.Context) error {
	restPresenter := rest.NewRestApiPresenter(ctx)

	var postBody PostAccountsRequest
	err := ctx.Bind(&postBody)
	if err != nil {
		restPresenter.SendError(
			http.StatusBadRequest,
			http.StatusText(http.StatusBadRequest),
			"Invalid format for createAccount",
		)
		return restPresenter.Error()
	}

	b.accounts.CreateAccount(restPresenter, postBody.Login, postBody.Password)
	return restPresenter.Error()
}
