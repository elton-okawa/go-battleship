package e2e

import (
	"fmt"
	"os"
	"path/filepath"
)

var files = [...]string{"accounts", "games"}

func CleanupDatabase(basePath string) {
	fmt.Printf("%s\n", basePath)
	for _, file := range files {
		path := filepath.Join(basePath, fmt.Sprintf("%s.json", file))

		if _, err := os.Stat(path); err == nil {
			if removeErr := os.Remove(path); removeErr != nil {
				panic(removeErr)
			}

		}
	}
}
