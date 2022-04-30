package presenter

import (
	"elton-okawa/battleship/internal/entity/gamestate"
)

type CommandLinePresenter struct {
}

// TODO implemente cmd presenter
func (cmd *CommandLinePresenter) StartResult(*gamestate.GameState, error) {
}

func (cmd *CommandLinePresenter) ShootResult(*gamestate.GameState, bool, int, error) {
}
