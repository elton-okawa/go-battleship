package controlaccount

import (
	"elton-okawa/battleship/internal/usecase/ucaccount"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	useCase ucaccount.UseCase
}

func New(a ucaccount.UseCase) Controller {
	return Controller{
		useCase: a,
	}
}

func (c Controller) CreateAccount(res ucaccount.Output, ctx echo.Context) error {
	var body CreateAccountJSONBody
	if err := ctx.Bind(&body); err != nil {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			"Invalid format for createAccount",
		)
	}

	return c.useCase.CreateAccount(res, body.Login, body.Password)
}

func (c Controller) Login(res ucaccount.Output, ctx echo.Context) error {
	var body AccountLoginJSONBody
	if err := ctx.Bind(&body); err != nil {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			"Invalid body to perform login",
		)
	}

	return c.useCase.Login(res, body.Login, body.Password)
}
