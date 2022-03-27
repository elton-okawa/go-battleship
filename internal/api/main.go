// Routing based on Axel Wagner's ShiftPath approach:
// https://blog.merovius.de/2017/06/18/how-not-to-use-an-http-router.html

package api

import (
	"elton-okawa/battleship/internal/api/games"
	"elton-okawa/battleship/internal/api/router"
	"elton-okawa/battleship/internal/database"
	"elton-okawa/battleship/internal/interface_adapter/controller"
	"elton-okawa/battleship/internal/interface_adapter/presenter"
	"elton-okawa/battleship/internal/use_case"
	"net/http"
)

type App struct {
	routers map[string]router.Router
}

func Init() *App {
	gameHandler := use_case.NewGame(
		&database.DefaultDatabase,
	)

	gamesController := controller.NewGamesController(
		gameHandler,
	)

	return &App{
		routers: map[string]router.Router{
			"games": games.NewGamesRouter(gamesController),
		},
	}
}

func (app *App) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	var resource string
	resource, r.URL.Path = router.ShiftPath(r.URL.Path)

	presenter := presenter.NewRestApiPresenter(rw)
	if router, exist := app.routers[resource]; exist {
		router.Route(presenter, r)
	} else {
		presenter.Error("Not Implemented", http.StatusNotImplemented)
	}
}
