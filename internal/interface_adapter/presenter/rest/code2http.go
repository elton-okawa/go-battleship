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
	errors.GenericError: {
		title: http.StatusText(http.StatusInternalServerError),
		code:  http.StatusInternalServerError,
	},
	errors.ElementNotFound: {
		title: http.StatusText(http.StatusNotFound),
		code:  http.StatusNotFound,
	},

	errors.IncorrectPassword: {
		title: http.StatusText(http.StatusUnauthorized),
		code:  http.StatusUnauthorized,
	},
}
