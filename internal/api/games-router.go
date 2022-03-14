package api

import (
	"elton-okawa/battleship/internal/interface_adapter/controller"
	"net/http"
)

type gamesRouter struct {
}

func (g *gamesRouter) route(rw http.ResponseWriter, r *http.Request) {
	var id string
	id, r.URL.Path = shiftPath(r.URL.Path)

	if id == "" {
		(&gamesHandler{id: id}).handle(rw, r)
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
	id string
}

func (gh *gamesHandler) handle(rw http.ResponseWriter, r *http.Request) {
	if handler, exist := gamesMethods[r.Method]; exist {
		handler(rw, r)
	} else {
		http.Error(rw, "Games method not allowed", http.StatusMethodNotAllowed)
	}
}

var gamesMethods map[string]handle = map[string]handle{
	"POST": postGames,
}

func postGames(res http.ResponseWriter, req *http.Request) {
	game := controller.PostGame()

	res.Write([]byte(game.Board.String()))
}
