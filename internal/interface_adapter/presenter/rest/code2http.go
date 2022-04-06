package rest

import (
	"elton-okawa/battleship/internal/use_case"
	"net/http"
)

type HttpError struct {
	title string
	code  int
}

var CodeToHttp = map[int]HttpError{
	use_case.ElementNotFound: HttpError{
		title: http.StatusText(http.StatusNotFound),
		code:  http.StatusNotFound,
	},
}
