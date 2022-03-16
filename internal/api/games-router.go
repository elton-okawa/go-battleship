package api

import (
	"elton-okawa/battleship/internal/interface_adapter/controller"
	"elton-okawa/battleship/internal/interface_adapter/presenter"
	"net/http"
)

type newPresenterFunc func(http.ResponseWriter) *presenter.RestApiPresenter

func NewGamesRouter(gc *controller.GamesController, np newPresenterFunc) *GamesRouter {
	return &GamesRouter{
		controller:   gc,
		newPresenter: np,
	}
}

type GamesRouter struct {
	controller   *controller.GamesController
	newPresenter newPresenterFunc
}

func (g *GamesRouter) route(rw http.ResponseWriter, r *http.Request) {
	var id string
	id, r.URL.Path = shiftPath(r.URL.Path)

	if id == "" {
		gh := &gamesHandler{
			id:         id,
			controller: g.controller,
			presenter:  g.newPresenter(rw),
		}
		gh.handle(rw, r)
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
	presenter  *presenter.RestApiPresenter
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
	gh.controller.PostGame(gh.presenter)

	// TODO como lidar com o presenter sendo chamado pelo use case e mandar resposta via http?
	// res.Write([]byte(game.Board.String()))
}
