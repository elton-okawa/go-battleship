package api

import (
	"elton-okawa/battleship/internal/entity"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type handleShoot func(http.ResponseWriter, *http.Request, string)

type gameActionShootRouter struct {
	gameId string
}

func prepareGameActionShootRouter(gameId string) router {
	return &gameActionShootRouter{
		gameId: gameId,
	}
}

func (sr *gameActionShootRouter) route(rw http.ResponseWriter, r *http.Request) {
	if handle, exist := shootMethods[r.Method]; exist {
		handle(rw, r, sr.gameId)
	} else {
		http.Error(rw, "Shoot action method not allowed", http.StatusMethodNotAllowed)
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

func postShoot(res http.ResponseWriter, req *http.Request, gameId string) {
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(res, "Invalid body", 500)
		return
	}

	// Unmarshal
	var body shootBody
	err = json.Unmarshal(data, &body)
	if err != nil {
		http.Error(res, "Body does not contain required fields", 500)
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
