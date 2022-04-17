package rest

import (
	"elton-okawa/battleship/internal/entity/account"
	"net/http"
)

type newAccountResponse struct {
	Id    string `json:"id"`
	Login string `json:"login"`
}

func (rp RestApiPresenter) CreateAccountResponse(e account.Account, err error) {
	if err != nil {
		rp.HandleError(err)
		return
	}

	res := newAccountResponse{
		Id:    e.Id,
		Login: e.Login,
	}

	rp.responseBody(http.StatusCreated, res)
}

type loginResponse struct {
	Token     string `json:"token"`
	ExpiresAt int64  `json:"expiresAt"`
}

func (rp RestApiPresenter) LoginResponse(e account.Account, token string, expires int64, err error) {
	if err != nil {
		rp.HandleError(err)
		return
	}

	res := loginResponse{
		Token:     token,
		ExpiresAt: expires,
	}

	rp.responseBody(http.StatusOK, res)
}
