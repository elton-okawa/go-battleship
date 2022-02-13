package main

import (
	"elton-okawa/battleship/internal/api"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/games", api.Games)

	fmt.Println("Server listening to :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
