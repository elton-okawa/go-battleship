package api

import (
	"elton-okawa/battleship/internal/interface_adapter/controller"
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

func (g *GamesRouter) route(rw http.ResponseWriter, r *http.Request) {
	var id string
	id, r.URL.Path = shiftPath(r.URL.Path)

	if id == "" {
		(&gamesHandler{id: id, controller: g.controller}).handle(rw, r)
	} else {
		var resource string
		resource, r.URL.Path = shiftPath(r.URL.Path)

		if router, exist := gamesSubRouters[resource]; exist {
			router(id).route(rw, r)
		} else {
			http.Error(rw, "Games resource not implemented", http.StatusNotImplemented)
		}
	}
}

var gamesSubRouters map[string]func(string) router = map[string]func(string) router{
	"actions": prepareGameActionRouter,
}

type gamesHandler struct {
	controller *controller.GamesController
	id         string
}

func (gh *gamesHandler) handle(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		gh.postGames(rw, r)
	default:
		http.Error(rw, "Games method not allowed", http.StatusMethodNotAllowed)
	}
}

func (gh *gamesHandler) postGames(rw http.ResponseWriter, r *http.Request) {
	gh.controller.PostGame()

	// TODO how to deal with it
	// res.Write([]byte(game.Board.String()))
}
