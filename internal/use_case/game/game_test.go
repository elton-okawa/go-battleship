package game

import (
	"testing"
)

// TODO better test
// TODO better mock
type MockPersistence struct {
}

func (mp *MockPersistence) SaveGameState(gs *GameState) error {
	return nil
}

func (mp *MockPersistence) GetGameState(id string) (*GameState, error) {
	return nil, nil
}

type MockOutput struct{}

func (mo *MockOutput) StartResult(*GameState, error) {
}

func (mo *MockOutput) ShootResult(*GameState, bool, int, error) {
}

func TestStartGame(t *testing.T) {
	game := Game{
		persistence: &MockPersistence{},
	}
	game.Start(&MockOutput{})
}
