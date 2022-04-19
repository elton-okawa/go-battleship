package rest

import (
	"elton-okawa/battleship/internal/entity"
	"elton-okawa/battleship/internal/usecase/game"
	"encoding/json"
	"fmt"
	"net/http"
)

type shootResponse struct {
	Hit   bool         `json:"hit"`
	Ships int          `json:"ships"`
	Board entity.Board `json:"board"`
}

func (rp RestApiPresenter) StartResult(gs *game.GameState, err error) {
	if err != nil {
		rp.MapError(err)
		return
	}

	// TODO map to a real response
	rp.responseBody(http.StatusCreated, gs.Board.String())
}

func (rp RestApiPresenter) ShootResult(gs *game.GameState, hit bool, ships int, err error) {
	if err != nil {
		rp.MapError(err)
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
