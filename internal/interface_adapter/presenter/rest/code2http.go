package rest

import (
	"elton-okawa/battleship/internal/use_case/errors"
	"net/http"
)

type HttpError struct {
	title string
	code  int
}

var CodeToHttp = map[int]HttpError{
	errors.ElementNotFound: HttpError{
		title: http.StatusText(http.StatusNotFound),
		code:  http.StatusNotFound,
	},
}
