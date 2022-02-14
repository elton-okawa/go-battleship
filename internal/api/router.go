// Routing based on Axel Wagner's ShiftPath approach:
// https://blog.merovius.de/2017/06/18/how-not-to-use-an-http-router.html

package api

import (
	"fmt"
	"net/http"
	"path"
	"strings"
)

type App struct {
	GameRouter *GameRouter
}

func (app *App) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	var resource string
	resource, req.URL.Path = ShiftPath(req.URL.Path)

	if resource == "games" {
		app.GameRouter.ServeHTTP(res, req)
	} else {
		http.Error(res, "Not Implemented", http.StatusNotImplemented)
	}
}

type GameRouter struct {
}

func (g *GameRouter) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	var resource string
	resource, req.URL.Path = ShiftPath(req.URL.Path)

	fmt.Println(resource)

	if req.URL.Path == "/" {
		switch req.Method {
		case "POST":
			g.handlePost(res, req)
		}
	}
}

func (g *GameRouter) handlePost(res http.ResponseWriter, req *http.Request) {
	game := PostGame()

	res.Write([]byte(game.Board.String()))
}

// Splits given path into <head>/<tail>
// Example: /users/10/receipts -> users, /10/receipts
func ShiftPath(p string) (head, tail string) {
	p = path.Clean("/" + p)
	i := strings.Index(p[1:], "/") + 1
	if i <= 0 {
		return p[1:], "/"
	}
	return p[1:i], p[i:]
}
