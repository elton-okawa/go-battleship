package main

type Start struct {
	// persistence game.GameStatePersistence
	// presenter   game.GameOutputBoundary
}

func (s *Start) Description() string {
	return "- start a game"
}

func (s *Start) Parse([]string) error {
	return nil
}

func (s *Start) Execute() (bool, error) {
	// game := game.NewGameUseCase(s.persistence)
	// game.Start(s.presenter)

	return false, nil
}
