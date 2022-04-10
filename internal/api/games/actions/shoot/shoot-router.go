package shoot

import (
	"elton-okawa/battleship/internal/api/router"
	"elton-okawa/battleship/internal/interface_adapter/controller"
	"elton-okawa/battleship/internal/interface_adapter/presenter/rest"
	"net/http"
)

type shootRouter struct {
	gameId     string
	controller *controller.GamesController
}

func NewShootRouter(controller *controller.GamesController, gameId string) router.Router {
	return &shootRouter{
		gameId:     gameId,
		controller: controller,
	}
}

func (sr *shootRouter) Route(p rest.RestApiPresenter, r *http.Request) {
	var head string
	head, r.URL.Path = router.ShiftPath(r.URL.Path)

	if head == "" {
		handler := &shootHandler{
			controller: sr.controller,
			gameId:     sr.gameId,
		}
		handler.handle(p, r)
	} else {
		p.Error("Not implemented", http.StatusNotImplemented)
	}
}
