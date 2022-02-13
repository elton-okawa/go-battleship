package api

import (
	"elton-okawa/battleship/internal/engine"
	"fmt"
	"net/http"
)

func Games(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		path := r.URL.Path[len("/games/"):]
		getGames(path)
	case "POST":
		v := postGames()
		fmt.Fprintf(w, "%s", v.Board)
	}
}

func getGames(id string) {
	fmt.Println("Get games")
}

func postGames() engine.GameState {
	return engine.StartGame()
}
