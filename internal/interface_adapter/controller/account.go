package controller

import "elton-okawa/battleship/internal/usecase/account"

type AccountController struct {
	useCase account.AccountUseCase
}

func NewAccountController(a account.AccountUseCase) AccountController {
	return AccountController{
		useCase: a,
	}
}

func (c AccountController) CreateAccount(res account.AccountOutput, login, password string) {
	c.useCase.CreateAccount(res, login, password)
}

func (c AccountController) Login(res account.AccountOutput, login, password string) {
	c.useCase.Login(res, login, password)
}
