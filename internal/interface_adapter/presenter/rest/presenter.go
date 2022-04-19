package rest

import (
	"elton-okawa/battleship/internal/usecase/ucerror"
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ProblemJson struct {
	Title      string `json:"title"`
	Status     int    `json:"status"`
	Detail     string `json:"detail"`
	Stacktrace string `json:"stack,omitempty"`
	// instance string
}

type ResponseCallback func(int, interface{})

type RestApiPresenter struct {
	code int
	body interface{}
	err  error
	// context echo.Context
	// err     error
}

func New() *RestApiPresenter {
	return &RestApiPresenter{}
}

func (rp *RestApiPresenter) responseBody(code int, data interface{}) {
	// rp.err = rp.context.JSON(code, data)
	rp.code = code
	rp.body = data
}

func (rp *RestApiPresenter) response(code int) {
	// rp.err = rp.context.NoContent(code)
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

	// rp.context.Response().Header().Set("Content-Type", "application/problem+json")
	// rp.err = rp.context.JSON(c, &p)

	return c, p
}

func (rp *RestApiPresenter) CreateError(code int, message string) {
	p := ProblemJson{
		Title:  http.StatusText(code),
		Status: code,
		Detail: message,
	}

	// rp.context.Response().Header().Set("Content-Type", "application/problem+json")
	// rp.err = rp.context.JSON(code, &p)
	rp.code = code
	rp.body = p
}

func (rp *RestApiPresenter) Error() error {
	return rp.err
}

func (rp *RestApiPresenter) Body() interface{} {
	return rp.body
}

func (rp *RestApiPresenter) Code() int {
	return rp.code
}
