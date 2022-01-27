package main

import (
	"elton-okawa/battleship/internal/engine"
)

type Start struct {
}

func (s Start) Parse([]string) error {
	return nil
}

func (s Start) Execute() {
	engine.StartGame()
}
