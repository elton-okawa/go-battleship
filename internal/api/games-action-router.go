package api

import (
	"net/http"
)

type gameActionsRouter struct {
	gameId string
}

func prepareGameActionRouter(gameId string) router {
	return &gameActionsRouter{
		gameId: gameId,
	}
}

func (ac *gameActionsRouter) route(rw http.ResponseWriter, r *http.Request) {
	var action string
	action, r.URL.Path = shiftPath(r.URL.Path)

	if router, exist := gameActionsSubRouter[action]; exist {
		router(ac.gameId).route(rw, r)
	} else {
		http.Error(rw, "Game action not implemented", http.StatusNotImplemented)
	}
}

var gameActionsSubRouter map[string]func(string) router = map[string]func(string) router{
	"shoot": prepareGameActionShootRouter,
}
