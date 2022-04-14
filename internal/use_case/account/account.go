package account

import (
	"elton-okawa/battleship/internal/entity"
	"elton-okawa/battleship/internal/use_case/errors"
	"fmt"
)

type AccountPersistence interface {
	SaveAccount(entity.Account) error
	GetAccount(login string) (entity.Account, error)
}

type AccountOutput interface {
	CreateAccountResponse(entity.Account, error)
	LoginResponse(entity.Account, string, int64, error)
}

type AccountUseCase struct {
	persistence AccountPersistence
}

func NewAccountUseCase(persistence AccountPersistence) AccountUseCase {
	return AccountUseCase{
		persistence: persistence,
	}
}

func (a AccountUseCase) CreateAccount(res AccountOutput, login, password string) {
	acc, err := entity.NewAccount(login, password)

	if err != nil {
		useCaseError := errors.NewError(
			fmt.Sprintf("Failed to create a new account for '%s'", login),
			errors.GenericError,
			err,
		)

		res.CreateAccountResponse(entity.Account{}, useCaseError)
		return
	}

	if a.persistence.SaveAccount(acc) != nil {
		useCaseError := errors.NewError(
			fmt.Sprintf("Failed to save a new account for '%s'", login),
			errors.GenericError,
			err,
		)

		res.CreateAccountResponse(entity.Account{}, useCaseError)
		return
	}

	res.CreateAccountResponse(acc, nil)
}

func (a AccountUseCase) Login(res AccountOutput, login, password string) {
	acc, err := a.persistence.GetAccount(login)

	if err != nil {
		useCaseError := errors.NewError(
			fmt.Sprintf("Account '%s' not found", login),
			errors.ElementNotFound,
			err,
		)
		res.LoginResponse(entity.Account{}, "", 0, useCaseError)
		return
	}

	if err = acc.Authenticate(password); err != nil {
		useCaseError := errors.NewError(
			"Incorrect password",
			errors.IncorrectPassword,
			err,
		)

		res.LoginResponse(entity.Account{}, "", 0, useCaseError)
		return
	}

	if token, expires, err := entity.NewJwtToken(login); err == nil {
		res.LoginResponse(acc, token, expires, err)
	} else {
		useCaseError := errors.NewError(
			"Error while creating JWT Token",
			errors.GenericError,
			err,
		)
		res.LoginResponse(entity.Account{}, "", 0, useCaseError)
	}
}
