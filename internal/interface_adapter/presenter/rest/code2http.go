package rest

import (
	"elton-okawa/battleship/internal/usecase/ucerror"
	"net/http"
)

type HttpError struct {
	title string
	code  int
}

var CodeToHttp = map[int]HttpError{
	ucerror.GenericError: {
		title: http.StatusText(http.StatusInternalServerError),
		code:  http.StatusInternalServerError,
	},
	ucerror.ElementNotFound: {
		title: http.StatusText(http.StatusNotFound),
		code:  http.StatusNotFound,
	},

	ucerror.IncorrectPassword: {
		title: http.StatusText(http.StatusUnauthorized),
		code:  http.StatusUnauthorized,
	},
}
