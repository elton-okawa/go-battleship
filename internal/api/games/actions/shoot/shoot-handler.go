package shoot

import (
	"elton-okawa/battleship/internal/interface_adapter/controller"
	"elton-okawa/battleship/internal/interface_adapter/presenter/rest"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type shootHandler struct {
	controller *controller.GamesController
	gameId     string
}

func (sh *shootHandler) handle(p *rest.RestApiPresenter, r *http.Request) {
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

func (sh *shootHandler) postShoot(p *rest.RestApiPresenter, r *http.Request) {
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
