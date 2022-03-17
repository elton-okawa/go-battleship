package api

import (
	"elton-okawa/battleship/internal/interface_adapter/controller"
	"elton-okawa/battleship/internal/interface_adapter/presenter"
	"net/http"
)

type gameActionsRouter struct {
	gameId     string
	controller *controller.GamesController
}

func newGameActionRouter(
	controller *controller.GamesController,
	gameId string,
) router {
	return &gameActionsRouter{
		gameId:     gameId,
		controller: controller,
	}
}

func (ac *gameActionsRouter) route(p *presenter.RestApiPresenter, r *http.Request) {
	var action string
	action, r.URL.Path = shiftPath(r.URL.Path)

	if router, exist := gameActionsSubRouter[action]; exist {
		router(ac.gameId).route(p, r)
	} else {
		p.Error("Game action not implemented", http.StatusNotImplemented)
	}
}

var gameActionsSubRouter map[string]func(string) router = map[string]func(string) router{
	"shoot": newGameActionShootRouter,
}
