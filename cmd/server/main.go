package main

import (
	"elton-okawa/battleship/internal/infra/database"
	"elton-okawa/battleship/internal/infra/router"
	"fmt"
	"math/rand"
	"path/filepath"
	"time"
)

var address = "localhost:8080"

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Printf("Server listening to %s\n", address)
	path, _ := filepath.Abs(filepath.Join("internal", "filedb"))
	opt := router.Options{
		Repo: database.RepositoryOption{
			Path: path,
		},
	}
	app, _ := router.Setup(opt)

	app.Logger.Fatal(app.Start(address))
}
