package ucgame

import (
	"elton-okawa/battleship/internal/entity/gamestate"
	"testing"
)

// TODO better test
// TODO better mock
type MockPersistence struct {
}

func (mp *MockPersistence) SaveGameState(gs *gamestate.GameState) error {
	return nil
}

func (mp *MockPersistence) GetGameState(id string) (*gamestate.GameState, error) {
	return nil, nil
}

type MockOutput struct{}

func (mo *MockOutput) StartResult(*gamestate.GameState, error) {
}

func (mo *MockOutput) ShootResult(*gamestate.GameState, bool, int, error) {
}

func TestStartGame(t *testing.T) {
	// game := UseCase{
	// 	persistence: &MockPersistence{},
	// }
	// game.Start(&MockOutput{})
}
