package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"elton-okawa/battleship/internal/database"
)

func main() {

	debugCommand := flag.NewFlagSet("debug", flag.ExitOnError)
	databaseCommand := flag.NewFlagSet("db", flag.ExitOnError)

	operationPtr := databaseCommand.String("operation", "", "Operation to perform in DB (Required)")
	idPtr := databaseCommand.String("id", "", "File id")
	fieldPtr := databaseCommand.String("field", "", "Simple field")

	if len(os.Args) <= 1 {
		fmt.Println("Please, provide a subcommand")
		flag.PrintDefaults()
		os.Exit(1)
	}

	command := os.Args[1]
	switch command {
	case "debug":
		debugCommand.Parse(os.Args[2:])
	case "db":
		databaseCommand.Parse(os.Args[2:])
	default:
		fmt.Println("Subcommand not found")
		flag.PrintDefaults()
		os.Exit(1)
	}

	if debugCommand.Parsed() {
		// use_case.StartGame()
	} else if databaseCommand.Parsed() {
		handleDbCommand(*operationPtr, *idPtr, *fieldPtr)
	}
}

type DebugModel struct {
	Id    string `json:"id"`
	Field string `json:"field"`
}

func handleDbCommand(operation string, id string, field string) {
	fmt.Printf("Db command (operation: %s, id: %s, field: %s)\n", operation, id, field)
	db := database.JsonDatabase{Filepath: "./db/index.json"}

	if operation == "" {
		fmt.Println("Operation must be passed")
		os.Exit(1)
	}

	switch operation {
	case "save":
		dm := DebugModel{Id: id, Field: field}
		if err := db.Save(dm); err != nil {
			log.Fatal(err)
		}
	case "get":
		var dm DebugModel
		err := db.Get(id, &dm)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Output: %+v\n", dm)
	}
}
