package account

import (
	"elton-okawa/battleship/internal/entity"
	"elton-okawa/battleship/internal/use_case/errors"
	"fmt"
)

type AccountPersistence interface {
	SaveAccount(entity.Account) error
}

type AccountOutput interface {
	CreateAccountResponse(entity.Account, error)
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
			errors.CreateAccountError,
			err,
		)

		res.CreateAccountResponse(acc, useCaseError)
		return
	}

	if a.persistence.SaveAccount(acc) != nil {
		useCaseError := errors.NewError(
			fmt.Sprintf("Failed to save a new account for '%s'", login),
			errors.CreateAccountError,
			err,
		)

		res.CreateAccountResponse(acc, useCaseError)
		return
	}

	res.CreateAccountResponse(acc, nil)
}
