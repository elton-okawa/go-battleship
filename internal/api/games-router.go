package api

import (
	"elton-okawa/battleship/internal/entity"
	"elton-okawa/battleship/internal/interface_adapter/controller"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type gamesRouter struct {
	handlers map[string]prepareStruct
	// methods   map[string]handle
	// resources map[string]handleSub
}

func initGamesRouter() *gamesRouter {
	return &gamesRouter{
		handlers: map[string]prepareStruct{
			"": prepareGame,
			// "actions": handleActions,
		},
		// methods: map[string]handle{
		// 	"POST": handlePost,
		// },
		// resources: map[string]handleSub{
		// 	"actions": handleActions,
		// },
	}
}

type prepareStruct func(string) routeHandler

var gameMethods map[string]handle = map[string]handle{
	"POST": postGames,
}

func prepareGame(head string) routeHandler {
	return &gameHandler{
		head:    head,
		methods: gameMethods,
	}
}

type routeHandler interface {
	handle(http.ResponseWriter, *http.Request)
}

type gameHandler struct {
	head    string
	methods map[string]handle
}

func (gh *gameHandler) handle(rw http.ResponseWriter, r *http.Request) {
	if handler, exist := gh.methods[r.Method]; exist {
		handler(rw, r)
	} else {
		http.Error(rw, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func postGames(res http.ResponseWriter, req *http.Request) {
	game := controller.PostGame()

	res.Write([]byte(game.Board.String()))
}

func (g *gamesRouter) handle(res http.ResponseWriter, req *http.Request) {
	var head string
	head, req.URL.Path = shiftPath(req.URL.Path)

	if handler, exist := g.handlers[head]; exist {
		handler(head).handle(res, req)
	} else {
		http.Error(res, "Method not allowed", http.StatusMethodNotAllowed)
	}

	// if req.URL.Path == "/" {

	// } else {
	// 	var resource string
	// 	resource, req.URL.Path = shiftPath(req.URL.Path)
	// 	if handleSub, exist := g.resources[resource]; exist {
	// 		handleSub(res, req, id)
	// 	} else {
	// 		http.Error(res, "Not implemented", http.StatusNotImplemented)
	// 	}
	// }
}

func handleActions(res http.ResponseWriter, req *http.Request, id string) {
	var action string
	action, req.URL.Path = shiftPath(req.URL.Path)

	switch action {
	case "shoot":
		handleShoot(res, req, id)
	default:
		http.Error(res, "Not implemented", http.StatusNotImplemented)
	}
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
