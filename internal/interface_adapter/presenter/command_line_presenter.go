package presenter

import (
	"elton-okawa/battleship/internal/ucerror/game"
)

type CommandLinePresenter struct {
}

// TODO implemente cmd presenter
func (cmd *CommandLinePresenter) StartResult(*game.GameState, error) {
}

func (cmd *CommandLinePresenter) ShootResult(*game.GameState, bool, int, error) {
}
