package accounts

import (
	"elton-okawa/battleship/internal/api/router"
	"elton-okawa/battleship/internal/interface_adapter/controller"
	"elton-okawa/battleship/internal/interface_adapter/presenter/rest"
	"net/http"
)

type AccountsRouter struct {
	controller controller.AccountController
}

func NewAccountsRouter(c controller.AccountController) AccountsRouter {
	return AccountsRouter{
		controller: c,
	}
}

func (a AccountsRouter) Route(p rest.RestApiPresenter, r *http.Request) {
	var id string
	id, r.URL.Path = router.ShiftPath(r.URL.Path)

	if id == "" {
		ah := accountsHandler{
			controller: a.controller,
		}
		ah.handle(p, r)
	} else {
		var resource string
		resource, r.URL.Path = router.ShiftPath(r.URL.Path)

		if router, exist := accountsSubRouters[resource]; exist {
			router(a.controller, id).Route(p, r)
		} else {
			p.Error("Games resource not implemented", http.StatusNotImplemented)
		}
	}
}

type accountsSubRouter func(controller.AccountController, string) router.Router

var accountsSubRouters map[string]accountsSubRouter = map[string]accountsSubRouter{
	// "actions": actions.NewActionsRouter,
}
