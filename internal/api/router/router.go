package router

import (
	"elton-okawa/battleship/internal/interface_adapter/presenter/rest"
	"net/http"
)

type Router interface {
	Route(*rest.RestApiPresenter, *http.Request)
}
