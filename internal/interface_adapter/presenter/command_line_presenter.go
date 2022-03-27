package presenter

import "elton-okawa/battleship/internal/use_case"

type CommandLinePresenter struct {
}

// TODO implemente cmd presenter
func (cmd *CommandLinePresenter) StartResult(*use_case.GameState, error) {
}

func (cmd *CommandLinePresenter) ShootResult(*use_case.GameState, bool, int, error) {
}
