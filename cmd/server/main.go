package main

import (
	"elton-okawa/battleship/internal/infra/router"
	"fmt"
)

var address = "localhost:8080"

func main() {
	fmt.Printf("Server listening to %s\n", address)
	app := router.SetupHandler()

	app.Logger.Fatal(app.Start(address))
}
