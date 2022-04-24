package setup

import (
	"elton-okawa/battleship/internal/infra/router"
	"fmt"
	"net/http/httptest"
	"os"
	"path/filepath"
)

var basePath, _ = filepath.Abs(filepath.Join("..", "..", "..", "db", "test"))
var files = [...]string{"accounts", "games"}

func CleanupDatabase() {
	fmt.Printf("%s\n", basePath)
	for _, file := range files {
		path := filepath.Join(basePath, fmt.Sprintf("%s.json", file))

		if _, err := os.Stat(path); err == nil {
			if removeErr := os.Remove(path); removeErr != nil {
				panic(removeErr)
			}

		}
	}
}

func TestServer() (*httptest.Server, *router.Repository) {
	opt := router.Options{
		Repo: router.RepositoryOption{
			Path: basePath,
		},
	}

	rt, db := router.Setup(opt)
	svr := httptest.NewServer(rt)
	fmt.Printf("Test server listening to '%s'\n", svr.URL)
	return svr, db
}
