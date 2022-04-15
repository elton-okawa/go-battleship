package controlaccount

import (
	"elton-okawa/battleship/internal/usecase/ucaccount"
)

type Controller struct {
	useCase ucaccount.UseCase
}

func New(a ucaccount.UseCase) Controller {
	return Controller{
		useCase: a,
	}
}

func (c Controller) CreateAccount(res ucaccount.Output, login, password string) {
	c.useCase.CreateAccount(res, login, password)
}

func (c Controller) Login(res ucaccount.Output, login, password string) {
	c.useCase.Login(res, login, password)
}
