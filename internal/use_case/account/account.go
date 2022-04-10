package account

import (
	"elton-okawa/battleship/internal/entity"
	"elton-okawa/battleship/internal/use_case/errors"
	"fmt"
)

type AccountPersistence interface {
	SavePlayer(entity.Player) error
}

type AccountOutput interface {
	CreateAccountResponse(entity.Player, error)
}

type AccountUseCase struct {
	persistence AccountPersistence
}

func NewAccountUseCase(persistence AccountPersistence) AccountUseCase {
	return AccountUseCase{
		persistence: persistence,
	}
}

// TODO we might want to impose some password restrictions like length, characters
func (a AccountUseCase) CreateAccount(res AccountOutput, login, password string) {
	player, err := entity.NewPlayer(login, password)

	if err != nil {
		useCaseError := errors.NewError(
			fmt.Sprintf("Failed to create a new account for '%s'", login),
			errors.CreateAccountError,
			err,
		)

		res.CreateAccountResponse(player, useCaseError)
		return
	}

	if a.persistence.SavePlayer(player) != nil {
		useCaseError := errors.NewError(
			fmt.Sprintf("Failed to save a new account for '%s'", login),
			errors.CreateAccountError,
			err,
		)

		res.CreateAccountResponse(player, useCaseError)
		return
	}

	res.CreateAccountResponse(player, nil)
}
