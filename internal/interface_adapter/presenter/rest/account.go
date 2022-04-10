package rest

import (
	"elton-okawa/battleship/internal/entity"
	"encoding/json"
	"net/http"
)

type newAccountResponse struct {
	Id    string `json:"id"`
	Login string `json:"login"`
}

func (rp RestApiPresenter) CreateAccountResponse(p entity.Player, err error) {
	if err != nil {
		rp.handleError(err)
		return
	}

	res := newAccountResponse{
		Id:    p.Id,
		Login: p.Login,
	}

	body, _ := json.Marshal(res)
	rp.responseBody(http.StatusCreated, body)
}
