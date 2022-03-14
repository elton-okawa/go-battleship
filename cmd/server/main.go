package main

import (
	"elton-okawa/battleship/internal/api"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Server listening to :8080")
	app := api.Init()
	log.Fatal(http.ListenAndServe(":8080", app))
}
