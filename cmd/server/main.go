package main

import (
	Api "elton-okawa/battleship/internal/api"
	"fmt"
)

var address = "localhost:8080"

func main() {
	fmt.Printf("Server listening to %s\n", address)
	app := Api.SetupHandler()

	app.Logger.Fatal(app.Start(address))
}
