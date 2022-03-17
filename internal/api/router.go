// Routing based on Axel Wagner's ShiftPath approach:
// https://blog.merovius.de/2017/06/18/how-not-to-use-an-http-router.html

package api

import (
	"elton-okawa/battleship/internal/database"
	"elton-okawa/battleship/internal/interface_adapter/controller"
	"elton-okawa/battleship/internal/interface_adapter/presenter"
	"elton-okawa/battleship/internal/use_case"
	"net/http"
	"path"
	"strings"
)

type router interface {
	route(*presenter.RestApiPresenter, *http.Request)
}

type handle func(presenter.RestApiPresenter, *http.Request)

type App struct {
	routers map[string]router
}

func Init() *App {
	gameHandler := use_case.NewGame(
		&database.DefaultDatabase,
	)

	gamesController := controller.NewGamesController(
		gameHandler,
	)

	return &App{
		routers: map[string]router{
			"games": NewGamesRouter(gamesController),
		},
	}
}

func (app *App) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	var resource string
	resource, r.URL.Path = shiftPath(r.URL.Path)

	presenter := presenter.NewRestApiPresenter(rw)
	if router, exist := app.routers[resource]; exist {
		router.route(presenter, r)
	} else {
		presenter.Error("Not Implemented", http.StatusNotImplemented)
	}
}

// Splits given path into <head>/<tail>
// Example - /users
// - /users -> users, /
// - / -> "", /
// Example - /users/10
// - /users/10 -> users, /10
// - /10 -> 10, /
// Example - /users/10/receipts
// - /users/10/receipts -> users, /10/receipts
// - /10/receipts -> 10, /receipts
// - /receipts -> receipts, /
//
func shiftPath(p string) (head, tail string) {
	p = path.Clean("/" + p)
	i := strings.Index(p[1:], "/") + 1
	if i <= 0 {
		return p[1:], "/"
	}
	return p[1:i], p[i:]
}
