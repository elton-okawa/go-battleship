package api

import (
	"elton-okawa/battleship/internal/interface_adapter/controller"
	"elton-okawa/battleship/internal/interface_adapter/presenter"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type gameActionShootRouter struct {
	gameId     string
	controller *controller.GamesController
}

func newGameActionShootRouter(controller *controller.GamesController, gameId string) router {
	return &gameActionShootRouter{
		gameId:     gameId,
		controller: controller,
	}
}

func (sr *gameActionShootRouter) route(p *presenter.RestApiPresenter, r *http.Request) {
	var head string
	head, r.URL.Path = shiftPath(r.URL.Path)

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

type shootHandler struct {
	controller *controller.GamesController
	gameId     string
}

func (sh *shootHandler) handle(p *presenter.RestApiPresenter, r *http.Request) {
	switch r.Method {
	case "POST":
		sh.postShoot(p, r)
	default:
		p.Error("Shoot method not allowed", http.StatusMethodNotAllowed)
	}
}

type shootBody struct {
	Row int `json:"row"`
	Col int `json:"col"`
}

func (sh *shootHandler) postShoot(p *presenter.RestApiPresenter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		p.Error("Invalid body", http.StatusInternalServerError)
		return
	}

	// Unmarshal
	var body shootBody
	err = json.Unmarshal(data, &body)
	if err != nil {
		p.Error("Body does not contain required fields", http.StatusInternalServerError)
		return
	}

	sh.controller.Shoot(p, sh.gameId, body.Row, body.Col)
}
