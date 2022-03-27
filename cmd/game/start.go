package main

import (
	"elton-okawa/battleship/internal/use_case"
)

type Start struct {
	persistence use_case.GameStatePersistence
	presenter   use_case.GameOutputBoundary
}

func (s *Start) Description() string {
	return "- start a game"
}

func (s *Start) Parse([]string) error {
	return nil
}

func (s *Start) Execute() (bool, error) {
	game := use_case.NewGame(s.persistence)
	game.Start(s.presenter)

	return false, nil
}
