package games

import (
	"elton-okawa/battleship/internal/interface_adapter/controller"
	"elton-okawa/battleship/internal/interface_adapter/presenter"
	"net/http"
)

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
