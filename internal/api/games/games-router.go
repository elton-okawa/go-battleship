package games

import (
	"elton-okawa/battleship/internal/api/games/actions"
	"elton-okawa/battleship/internal/api/router"
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

func (g *GamesRouter) Route(p *presenter.RestApiPresenter, r *http.Request) {
	var id string
	id, r.URL.Path = router.ShiftPath(r.URL.Path)

	if id == "" {
		gh := &gamesHandler{
			id:         id,
			controller: g.controller,
		}
		gh.handle(p, r)
	} else {
		var resource string
		resource, r.URL.Path = router.ShiftPath(r.URL.Path)

		if router, exist := gamesSubRouters[resource]; exist {
			router(g.controller, id).Route(p, r)
		} else {
			p.Error("Games resource not implemented", http.StatusNotImplemented)
		}
	}
}

type gamesSubRouter func(*controller.GamesController, string) router.Router

var gamesSubRouters map[string]gamesSubRouter = map[string]gamesSubRouter{
	"actions": actions.NewActionsRouter,
}
