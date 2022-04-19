package rest

import (
	"elton-okawa/battleship/internal/usecase/ucaccount"
	"net/http"
)

type newAccountResponse struct {
	Id    string `json:"id"`
	Login string `json:"login"`
}

func (rp *RestApiPresenter) CreateAccountResponse(e ucaccount.AccountDTO) {
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

func (rp *RestApiPresenter) LoginResponse(dto ucaccount.LoginDTO) {
	res := loginResponse{
		Token:     dto.Token,
		ExpiresAt: dto.ExpiresAt,
	}

	rp.responseBody(http.StatusOK, res)
}
