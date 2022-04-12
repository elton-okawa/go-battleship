package rest

import (
	"elton-okawa/battleship/internal/entity"
	"net/http"
)

type newAccountResponse struct {
	Id    string `json:"id"`
	Login string `json:"login"`
}

func (rp RestApiPresenter) CreateAccountResponse(e entity.Account, err error) {
	if err != nil {
		rp.handleError(err)
		return
	}

	res := newAccountResponse{
		Id:    e.Id,
		Login: e.Login,
	}

	rp.responseBody(http.StatusCreated, res)
}
