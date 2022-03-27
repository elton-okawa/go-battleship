package rest

import (
	"elton-okawa/battleship/internal/entity"
	"elton-okawa/battleship/internal/use_case"
	"encoding/json"
	"fmt"
	"net/http"
)

type RestApiPresenter struct {
	responseWriter http.ResponseWriter
}

func NewRestApiPresenter(rw http.ResponseWriter) *RestApiPresenter {
	return &RestApiPresenter{
		responseWriter: rw,
	}
}

func (rp *RestApiPresenter) StartResult(gs *use_case.GameState, err error) {
	// TODO handle error
	rp.responseWriter.Write([]byte(gs.Board.String()))
}

type shootResponse struct {
	Hit   bool         `json:"hit"`
	Ships int          `json:"ships"`
	Board entity.Board `json:"board"`
}

func (rp *RestApiPresenter) ShootResult(gs *use_case.GameState, hit bool, ships int, err error) {

	if err != nil {
		fmt.Printf("%v", err)
		http.Error(rp.responseWriter, err.Error(), 400)
		return
	}

	shootRes := shootResponse{
		Hit:   hit,
		Ships: ships,
		Board: gs.Board,
	}

	fmt.Println(gs.Board.String())
	resData, _ := json.Marshal(shootRes)
	rp.responseWriter.Header().Set("Content-Type", "application/json")
	rp.responseWriter.Write(resData)
}

func (rp *RestApiPresenter) Error(message string, code int) {
	http.Error(rp.responseWriter, message, code)
}
