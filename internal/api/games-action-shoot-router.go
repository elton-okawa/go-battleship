package api

import (
	"elton-okawa/battleship/internal/entity"
	"elton-okawa/battleship/internal/interface_adapter/presenter"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type handleShoot func(*presenter.RestApiPresenter, *http.Request, string)

type gameActionShootRouter struct {
	gameId string
}

func newGameActionShootRouter(gameId string) router {
	return &gameActionShootRouter{
		gameId: gameId,
	}
}

func (sr *gameActionShootRouter) route(p *presenter.RestApiPresenter, r *http.Request) {
	if handle, exist := shootMethods[r.Method]; exist {
		handle(p, r, sr.gameId)
	} else {
		p.Error("Shoot action method not allowed", http.StatusMethodNotAllowed)
	}
}

var shootMethods map[string]handleShoot = map[string]handleShoot{
	"POST": postShoot,
}

type shootBody struct {
	Row int `json:"row"`
	Col int `json:"col"`
}

type shootResponse struct {
	Hit   bool         `json:"hit"`
	Ships int          `json:"ships"`
	Board entity.Board `json:"board"`
}

func postShoot(p *presenter.RestApiPresenter, r *http.Request, gameId string) {
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

	// hit, ships, gameState := controller.Shoot(gameId, body.Row, body.Col)
	// shootRes := shootResponse{
	// 	Hit:   hit,
	// 	Ships: ships,
	// 	Board: gameState.Board,
	// }

	// fmt.Println(gameState.Board.String())
	// resData, _ := json.Marshal(shootRes)
	// res.Header().Set("Content-Type", "application/json")
	// res.Write(resData)
}
