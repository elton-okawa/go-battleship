package ucaccount

import "elton-okawa/battleship/internal/entity/account"

type AccountDTO struct {
	Id    string
	Login string
}

func NewAccountDTO(e account.Account) AccountDTO {
	return AccountDTO{
		Id:    e.Id,
		Login: e.Login,
	}
}

type LoginDTO struct {
	Id        string
	Login     string
	Token     string
	ExpiresAt int64
}

func NewLoginDTO(e account.Account, token string, expiresAt int64) LoginDTO {
	return LoginDTO{
		Id:        e.Id,
		Login:     e.Login,
		Token:     token,
		ExpiresAt: expiresAt,
	}
}
