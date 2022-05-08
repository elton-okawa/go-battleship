package setup

import (
	"elton-okawa/battleship/internal/infra/database"
	"elton-okawa/battleship/internal/infra/router"
	"elton-okawa/battleship/internal/test"
	"fmt"
	"net/http/httptest"
)

func TestServer() (*httptest.Server, *router.Repository) {
	opt := router.Options{
		Repo: database.RepositoryOption{
			Path: test.DbFilePath(),
		},
	}

	rt, db := router.Setup(opt)
	svr := httptest.NewServer(rt)
	fmt.Printf("Test server listening to '%s'\n", svr.URL)
	return svr, db
}
