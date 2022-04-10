package rest

import (
	"elton-okawa/battleship/internal/entity"
	"elton-okawa/battleship/internal/use_case"
	"encoding/json"
	"fmt"
	"net/http"
)

type shootResponse struct {
	Hit   bool         `json:"hit"`
	Ships int          `json:"ships"`
	Board entity.Board `json:"board"`
}

func (rp RestApiPresenter) StartResult(gs *use_case.GameState, err error) {
	if err != nil {
		rp.handleError(err)
		return
	}

	rp.responseBody(http.StatusCreated, []byte(gs.Board.String()))
}

func (rp RestApiPresenter) ShootResult(gs *use_case.GameState, hit bool, ships int, err error) {
	if err != nil {
		rp.handleError(err)
		return
	}

	shootRes := shootResponse{
		Hit:   hit,
		Ships: ships,
		Board: gs.Board,
	}

	fmt.Println(gs.Board.String())
	resData, _ := json.Marshal(shootRes)
	rp.responseBody(http.StatusOK, resData)
}
