package use_case

type GameStatePersistence interface {
	SaveGameState(gs *GameState) error
	GetGameState(id string) (*GameState, error)
}
