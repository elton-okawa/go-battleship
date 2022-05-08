package test

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var files = [...]string{"accounts", "games", "game-state"}

func DbFilePath() string {
	dir, _ := os.Getwd()
	current := dir
	for !strings.HasSuffix(current, "internal") {
		current = filepath.Dir(current)
	}

	return filepath.Join(current, "filedb", "test")
}

func CleanupDatabase() {
	fmt.Printf("%s\n", DbFilePath())
	for _, file := range files {
		path := filepath.Join(DbFilePath(), fmt.Sprintf("%s.json", file))

		if _, err := os.Stat(path); err == nil {
			if removeErr := os.Remove(path); removeErr != nil {
				panic(removeErr)
			}

		}
	}
}
