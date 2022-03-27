package router

import (
	"elton-okawa/battleship/internal/interface_adapter/presenter"
	"net/http"
)

type Router interface {
	Route(*presenter.RestApiPresenter, *http.Request)
}
