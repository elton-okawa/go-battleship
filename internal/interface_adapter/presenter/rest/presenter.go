package rest

import (
	"elton-okawa/battleship/internal/entity"
	"elton-okawa/battleship/internal/use_case"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type ProblemJson struct {
	Title  string `json:"title"`
	Status int    `json:"status"`
	Detail string `json:"detail"`
	// instance string
}

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
	rp.responseWriter.Header().Set("Content-Type", "application/json")
	rp.responseWriter.Write(resData)
}

func (rp *RestApiPresenter) handleError(err error) {
	var e *use_case.UseCaseError
	var p ProblemJson
	var c int
	if errors.As(err, &e) {
		httpError := CodeToHttp[e.Code]
		c = httpError.code
		p = ProblemJson{
			Title:  httpError.title,
			Status: c,
			Detail: e.Message,
		}
	} else {
		c = http.StatusInternalServerError
		p = ProblemJson{
			Title:  http.StatusText(c),
			Status: c,
			Detail: fmt.Sprintf("An unexpected error occurred: %v", err),
		}
	}

	res, _ := json.Marshal(&p)
	rp.responseWriter.Header().Set("Content-Type", "application/problem+json")
	rp.responseWriter.WriteHeader(c)
	rp.responseWriter.Write(res)
}

func (rp *RestApiPresenter) Error(message string, code int) {
	http.Error(rp.responseWriter, message, code)
}
