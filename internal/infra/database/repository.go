package database

import (
	"fmt"
	"path/filepath"
)

type RepositoryOption struct {
	Path string
}

func (opt RepositoryOption) File(key string) string {
	path := filepath.Join(opt.Path, fmt.Sprintf("%s.json", key))
	fmt.Printf("[DB] %s: %s\n", key, path)
	return path
}
