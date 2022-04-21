package rest

import (
	"elton-okawa/battleship/internal/usecase/ucaccount"
	"net/http"
)

func (rp *RestApiPresenter) CreateAccountResponse(e ucaccount.AccountDTO) {
	res := PostAccountsResponse{
		Id:    e.Id,
		Login: e.Login,
	}

	rp.responseBody(http.StatusCreated, res)
}

func (rp *RestApiPresenter) LoginResponse(dto ucaccount.LoginDTO) {
	res := PostLoginResponse{
		Token:     dto.Token,
		ExpiresAt: dto.ExpiresAt,
	}

	rp.responseBody(http.StatusOK, res)
}
