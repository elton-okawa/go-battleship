package use_case

type GameStatePersistence interface {
	SaveGameState(gs *GameState)
	GetGameState() *GameState
}
