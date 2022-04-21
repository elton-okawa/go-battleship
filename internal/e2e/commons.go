package e2e

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

func SetupTestServer() *httptest.Server {
	opt := router.Options{
		Db: router.DBOptions{
			Path: basePath,
		},
	}

	rt := router.Setup(opt)
	svr := httptest.NewServer(rt)
	fmt.Printf("Test server listening to '%s'\n", svr.URL)
	return svr
}
