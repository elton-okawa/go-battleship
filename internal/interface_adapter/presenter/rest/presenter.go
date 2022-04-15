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

type RestApiPresenter struct {
	context echo.Context
	err     error
}

func NewRestApiPresenter(ctx echo.Context) *RestApiPresenter {
	return &RestApiPresenter{
		context: ctx,
	}
}

func (rp *RestApiPresenter) responseBody(code int, data interface{}) {
	rp.err = rp.context.JSON(code, data)
}

func (rp *RestApiPresenter) response(code int) {
	rp.err = rp.context.NoContent(code)
}

func (rp *RestApiPresenter) handleError(err error) {
	var e *ucerror.Error
	var p ProblemJson
	var c int

	fmt.Printf("%+w", err)
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

	rp.context.Response().Header().Set("Content-Type", "application/problem+json")
	rp.err = rp.context.JSON(c, &p)
}

func (rp *RestApiPresenter) SendError(code int, message string) {
	p := ProblemJson{
		Title:  http.StatusText(code),
		Status: code,
		Detail: message,
	}

	rp.context.Response().Header().Set("Content-Type", "application/problem+json")
	rp.err = rp.context.JSON(code, &p)
}

func (rp *RestApiPresenter) Error() error {
	return rp.err
}
