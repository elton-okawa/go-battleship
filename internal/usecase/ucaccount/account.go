package ucaccount

import (
	"elton-okawa/battleship/internal/entity/account"
	"elton-okawa/battleship/internal/entity/jwttoken"
	"elton-okawa/battleship/internal/usecase/ucerror"
	"fmt"
)

type Repository interface {
	Save(account.Account) error
	Get(login string) (account.Account, error)
}

type Output interface {
	CreateAccountResponse(AccountDTO)
	LoginResponse(LoginDTO)
}

type UseCase struct {
	repo Repository
}

func New(repo Repository) UseCase {
	return UseCase{
		repo: repo,
	}
}

func (a UseCase) CreateAccount(res Output, login, password string) error {
	acc, err := account.New(login, password)

	if err != nil {
		useCaseError := ucerror.New(
			fmt.Sprintf("Failed to create a new account for '%s'", login),
			ucerror.GenericError,
			err,
		)
		return useCaseError
	}

	if err = a.repo.Save(acc); err != nil {
		useCaseError := ucerror.New(
			fmt.Sprintf("Failed to save a new account for '%s'", login),
			ucerror.GenericError,
			err,
		)
		return useCaseError
	}

	res.CreateAccountResponse(NewAccountDTO(acc))
	return nil
}

func (a UseCase) Login(res Output, login, password string) error {
	acc, err := a.repo.Get(login)

	if err != nil {
		useCaseError := ucerror.New(
			fmt.Sprintf("Account '%s' not found", login),
			ucerror.IncorrectUsername,
			err,
		)
		return useCaseError
	}

	if err = acc.Authenticate(password); err != nil {
		useCaseError := ucerror.New(
			"Incorrect password",
			ucerror.IncorrectPassword,
			err,
		)
		return useCaseError
	}

	if token, expires, err := jwttoken.New(login, acc.PlayerId); err != nil {
		useCaseError := ucerror.New(
			"Error while creating JWT Token",
			ucerror.GenericError,
			err,
		)
		return useCaseError
	} else {
		res.LoginResponse(NewLoginDTO(acc, token, expires))
		return nil
	}
}
