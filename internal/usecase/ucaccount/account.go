package ucaccount

import (
	"elton-okawa/battleship/internal/entity/account"
	"elton-okawa/battleship/internal/entity/jwttoken"
	"elton-okawa/battleship/internal/usecase/ucerror"
	"fmt"
)

type Persistence interface {
	Save(account.Account) error
	Get(login string) (account.Account, error)
}

type Output interface {
	CreateAccountResponse(account.Account, error)
	LoginResponse(account.Account, string, int64, error)
}

type UseCase struct {
	persistence Persistence
}

func New(persistence Persistence) UseCase {
	return UseCase{
		persistence: persistence,
	}
}

func (a UseCase) CreateAccount(res Output, login, password string) {
	acc, err := account.New(login, password)

	if err != nil {
		useCaseError := ucerror.NewError(
			fmt.Sprintf("Failed to create a new account for '%s'", login),
			ucerror.GenericError,
			err,
		)

		res.CreateAccountResponse(account.Account{}, useCaseError)
		return
	}

	if a.persistence.Save(acc) != nil {
		useCaseError := ucerror.NewError(
			fmt.Sprintf("Failed to save a new account for '%s'", login),
			ucerror.GenericError,
			err,
		)

		res.CreateAccountResponse(account.Account{}, useCaseError)
		return
	}

	res.CreateAccountResponse(acc, nil)
}

func (a UseCase) Login(res Output, login, password string) {
	acc, err := a.persistence.Get(login)

	if err != nil {
		useCaseError := ucerror.NewError(
			fmt.Sprintf("Account '%s' not found", login),
			ucerror.ElementNotFound,
			err,
		)
		res.LoginResponse(account.Account{}, "", 0, useCaseError)
		return
	}

	if err = acc.Authenticate(password); err != nil {
		useCaseError := ucerror.NewError(
			"Incorrect password",
			ucerror.IncorrectPassword,
			nil,
		)

		res.LoginResponse(account.Account{}, "", 0, useCaseError)
		return
	}

	if token, expires, err := jwttoken.New(login); err == nil {
		res.LoginResponse(acc, token, expires, err)
	} else {
		useCaseError := ucerror.NewError(
			"Error while creating JWT Token",
			ucerror.GenericError,
			err,
		)
		res.LoginResponse(account.Account{}, "", 0, useCaseError)
	}
}
