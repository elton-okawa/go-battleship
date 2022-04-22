package main

import (
	"elton-okawa/battleship/internal/infra/router"
	"fmt"
	"path/filepath"
)

var address = "localhost:8080"

func main() {
	fmt.Printf("Server listening to %s\n", address)
	path, _ := filepath.Abs(filepath.Join("db"))
	opt := router.Options{
		Db: router.DBOptions{
			Path: path,
		},
	}
	app, _ := router.Setup(opt)

	app.Logger.Fatal(app.Start(address))
}
