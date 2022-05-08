package test

import (
	"fmt"
	"os"
	"path/filepath"
)

var BaseDBFilePath, _ = filepath.Abs(filepath.Join("..", "..", "..", "..", "db", "test"))
var files = [...]string{"accounts", "games", "game-state"}

func CleanupDatabase() {
	fmt.Printf("%s\n", BaseDBFilePath)
	for _, file := range files {
		path := filepath.Join(BaseDBFilePath, fmt.Sprintf("%s.json", file))

		if _, err := os.Stat(path); err == nil {
			if removeErr := os.Remove(path); removeErr != nil {
				panic(removeErr)
			}

		}
	}
}
