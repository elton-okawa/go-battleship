// Routing based on Axel Wagner's ShiftPath approach:
// https://blog.merovius.de/2017/06/18/how-not-to-use-an-http-router.html

package api

import (
	"net/http"
	"path"
	"strings"
)

type handler interface {
	handle(http.ResponseWriter, *http.Request)
}

type handle func(http.ResponseWriter, *http.Request)
type handleSub func(http.ResponseWriter, *http.Request, string)

type App struct {
	routers map[string]handler
}

func Init() *App {
	return &App{
		routers: map[string]handler{
			"games": initGamesRouter(),
		},
	}
}

func (app *App) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	var resource string
	resource, req.URL.Path = shiftPath(req.URL.Path)

	if router, exist := app.routers[resource]; exist {
		router.handle(res, req)
	} else {
		http.Error(res, "Not Implemented", http.StatusNotImplemented)
	}
}

// Splits given path into <head>/<tail>
// Example - /users
// - /users -> users, /
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
