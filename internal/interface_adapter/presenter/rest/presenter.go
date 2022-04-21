package rest

import (
	"elton-okawa/battleship/internal/usecase/ucerror"
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ProblemJson struct {
	Title  string `json:"title"`
	Status int    `json:"status"`
	Detail string `json:"detail"`
	Debug  string `json:"debug,omitempty"`
	// instance string
}

type RestApiPresenter struct {
	code int
	body interface{}
}

func New() *RestApiPresenter {
	return &RestApiPresenter{}
}

func (rp *RestApiPresenter) responseBody(code int, data interface{}) {
	rp.code = code
	rp.body = data
}

func (rp *RestApiPresenter) response(code int) {
	rp.code = code
}

func (rp *RestApiPresenter) MapError(err error) (int, interface{}) {
	var useCaseError *ucerror.Error
	var echoError *echo.HTTPError
	var p ProblemJson
	var c int

	if errors.As(err, &useCaseError) {
		httpError := CodeToHttp[useCaseError.Code]
		c = httpError.code

		// overwrite usecase message if necessary
		msg := httpError.message
		if msg == "" {
			msg = useCaseError.Message
		}

		p = ProblemJson{
			Title:  httpError.title,
			Status: c,
			Detail: msg,
			Debug:  useCaseError.Debug(), // TODO omit complete message on prod env
		}
	} else if errors.As(err, &echoError) {
		c = echoError.Code
		var msg = "no message"
		if v, ok := echoError.Message.(string); ok {
			msg = v
		}

		p = ProblemJson{
			Title:  http.StatusText(c),
			Status: c,
			Detail: msg,
		}
	} else {
		c = http.StatusInternalServerError
		p = ProblemJson{
			Title:  http.StatusText(c),
			Status: c,
			Detail: fmt.Sprintf("An unexpected error occurred: %v", err),
		}
	}

	return c, &p
}

func (rp *RestApiPresenter) Body() interface{} {
	return rp.body
}

func (rp *RestApiPresenter) Code() int {
	return rp.code
}
