package main

import (
	"elton-okawa/battleship/internal/use_case"
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
	use_case.StartGame()

	return false, nil
}
