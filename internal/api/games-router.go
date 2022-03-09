package api

import (
	"elton-okawa/battleship/internal/entity"
	"elton-okawa/battleship/internal/interface_adapter/controller"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type GamesRouter struct {
}

func (g *GamesRouter) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	var id string
	id, req.URL.Path = ShiftPath(req.URL.Path)

	if req.URL.Path == "/" {
		switch req.Method {
		case "POST":
			// TODO should handlePost receive res and req?
			g.handlePost(res, req)
		}
	} else {
		var resource string
		resource, req.URL.Path = ShiftPath(req.URL.Path)
		switch resource {
		case "actions":
			actionsHandler(id).ServeHTTP(res, req)
		default:
			http.Error(res, "Not Found", http.StatusNotFound)
		}
	}
}

func (g *GamesRouter) handlePost(res http.ResponseWriter, req *http.Request) {
	game := controller.PostGame()

	res.Write([]byte(game.Board.String()))
}

func actionsHandler(id string) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		var action string
		action, req.URL.Path = ShiftPath(req.URL.Path)

		switch action {
		case "shoot":
			handleShoot(res, req, id)
		default:
			http.Error(res, "Not implemented", http.StatusNotImplemented)
		}
	})
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

func handleShoot(res http.ResponseWriter, req *http.Request, id string) {

	switch req.Method {
	case "POST":
		// TODO should postShoot deal with res?
		postShoot(res, req, id)
	default:
		http.Error(res, "Not implemented", http.StatusNotImplemented)
	}
}

func postShoot(res http.ResponseWriter, req *http.Request, id string) {
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

	hit, ships, gameState := controller.Shoot(id, body.Row, body.Col)
	shootRes := shootResponse{
		Hit:   hit,
		Ships: ships,
		Board: gameState.Board,
	}

	fmt.Println(gameState.Board.String())
	resData, _ := json.Marshal(shootRes)
	res.Header().Set("Content-Type", "application/json")
	res.Write(resData)
}
