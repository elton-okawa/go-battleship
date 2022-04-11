package game

type GameStatePersistence interface {
	SaveGameState(gs *GameState) error
	GetGameState(id string) (*GameState, error)
}

type GameOutputBoundary interface {
	StartResult(*GameState, error)
	ShootResult(*GameState, bool, int, error)
}
