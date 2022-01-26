package main

import (
	"flag"
	"fmt"
	"os"

	"elton-okawa/battleship/internal/engine"
)

func main() {

	debugCommand := flag.NewFlagSet("debug", flag.ExitOnError)

	if len(os.Args) <= 1 {
		fmt.Println("Please, provide a subcommand")
		flag.PrintDefaults()
		os.Exit(1)
	}

	command := os.Args[1]
	switch command {
	case "debug":
		debugCommand.Parse(os.Args[2:])
	default:
		fmt.Println("Subcommand not found")
		flag.PrintDefaults()
		os.Exit(1)
	}

	if debugCommand.Parsed() {
		engine.StartGame()
	}
}
