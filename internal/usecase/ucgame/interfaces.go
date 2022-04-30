package ucgame

import "elton-okawa/battleship/internal/entity/gamestate"

type GameStateRepository interface {
	Save(gamestate.GameState) error
	Get(string) (*gamestate.GameState, error)
}

type GameOutputBoundary interface {
	StartResult(*gamestate.GameState, error)
	ShootResult(*gamestate.GameState, bool, int, error)
}
