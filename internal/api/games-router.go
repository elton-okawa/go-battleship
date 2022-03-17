package api

import (
	"elton-okawa/battleship/internal/interface_adapter/controller"
	"elton-okawa/battleship/internal/interface_adapter/presenter"
	"net/http"
)

func NewGamesRouter(gc *controller.GamesController) *GamesRouter {

	return &GamesRouter{
		controller: gc,
	}
}

type GamesRouter struct {
	controller *controller.GamesController
}

func (g *GamesRouter) route(p *presenter.RestApiPresenter, r *http.Request) {
	var id string
	id, r.URL.Path = shiftPath(r.URL.Path)

	if id == "" {
		gh := &gamesHandler{
			id:         id,
			controller: g.controller,
		}
		gh.handle(p, r)
	} else {
		var resource string
		resource, r.URL.Path = shiftPath(r.URL.Path)

		if router, exist := gamesSubRouters[resource]; exist {
			router(g.controller, id).route(p, r)
		} else {
			p.Error("Games resource not implemented", http.StatusNotImplemented)
		}
	}
}

type gamesSubRouter func(*controller.GamesController, string) router

var gamesSubRouters map[string]gamesSubRouter = map[string]gamesSubRouter{
	"actions": newGameActionRouter,
}

type gamesHandler struct {
	controller *controller.GamesController
	id         string
}

func (gh *gamesHandler) handle(p *presenter.RestApiPresenter, r *http.Request) {
	switch r.Method {
	case "POST":
		gh.postGames(p, r)
	default:
		p.Error("Games method not allowed", http.StatusMethodNotAllowed)
	}
}

func (gh *gamesHandler) postGames(p *presenter.RestApiPresenter, r *http.Request) {
	gh.controller.PostGame(p)
}
