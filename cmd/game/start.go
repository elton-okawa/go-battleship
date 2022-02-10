package main

import (
	"elton-okawa/battleship/internal/engine"
)

type Start struct {
}

func (s *Start) Description() string {
	return "- start a game"
}

func (s *Start) Parse([]string) error {
	return nil
}

func (s *Start) Execute() (bool, error) {
	engine.StartGame()

	return false, nil
}
