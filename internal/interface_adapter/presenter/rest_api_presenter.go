package presenter

import (
	"elton-okawa/battleship/internal/use_case"
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

func (rp *RestApiPresenter) ShootResult(gs *use_case.GameState, hit bool, ships int, err error) {
	fmt.Printf("hit: %t, ships: %d\n", hit, ships)
}

func (rp *RestApiPresenter) Error(message string, code int) {
	http.Error(rp.responseWriter, message, code)
}
