package rest

import (
	use_case_errors "elton-okawa/battleship/internal/use_case/errors"
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

func NewRestApiPresenter(rw http.ResponseWriter) RestApiPresenter {
	return RestApiPresenter{
		responseWriter: rw,
	}
}

func (rp RestApiPresenter) responseBody(code int, data []byte) {
	rp.responseWriter.Header().Set("Content-Type", "application/json")
	rp.responseWriter.WriteHeader(code)
	rp.responseWriter.Write(data)
}

func (rp RestApiPresenter) response(code int) {
	rp.responseWriter.Header().Set("Content-Type", "application/json")
	rp.responseWriter.WriteHeader(code)
}

func (rp RestApiPresenter) handleError(err error) {
	var e *use_case_errors.UseCaseError
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
