package accounts

import (
	"elton-okawa/battleship/internal/interface_adapter/controller"
	"elton-okawa/battleship/internal/interface_adapter/presenter/rest"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type accountsHandler struct {
	controller controller.AccountController
}

func newAccountsHandler(controller controller.AccountController) accountsHandler {
	return accountsHandler{
		controller: controller,
	}
}

func (ah accountsHandler) handle(p rest.RestApiPresenter, r *http.Request) {
	switch r.Method {
	case "POST":
		ah.postAccounts(p, r)
	default:
		p.Error("Accounts method not allowed", http.StatusMethodNotAllowed)
	}
}

type postAccountsBody struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (ah accountsHandler) postAccounts(p rest.RestApiPresenter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		p.Error("Error while reading body", http.StatusInternalServerError)
		return
	}

	var body postAccountsBody
	if json.Unmarshal(data, &body) != nil {
		p.Error("Invalid body", http.StatusBadRequest)
		return
	}

	ah.controller.CreateAccount(p, body.Login, body.Password)
}
