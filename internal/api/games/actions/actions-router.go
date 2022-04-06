package actions

import (
	"elton-okawa/battleship/internal/api/games/actions/shoot"
	"elton-okawa/battleship/internal/api/router"
	"elton-okawa/battleship/internal/interface_adapter/controller"
	"elton-okawa/battleship/internal/interface_adapter/presenter/rest"
	"net/http"
)

type actionsRouter struct {
	gameId     string
	controller *controller.GamesController
}

func NewActionsRouter(
	controller *controller.GamesController,
	gameId string,
) router.Router {
	return &actionsRouter{
		gameId:     gameId,
		controller: controller,
	}
}

func (ac *actionsRouter) Route(p *rest.RestApiPresenter, r *http.Request) {
	var action string
	action, r.URL.Path = router.ShiftPath(r.URL.Path)

	if router, exist := gameActionsSubRouters[action]; exist {
		router(ac.controller, ac.gameId).Route(p, r)
	} else {
		p.Error("Game action not implemented", http.StatusNotImplemented)
	}
}

type actionsSubRouter func(*controller.GamesController, string) router.Router

var gameActionsSubRouters map[string]actionsSubRouter = map[string]actionsSubRouter{
	"shoot": shoot.NewShootRouter,
}
