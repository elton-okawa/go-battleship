package use_case

import (
	"elton-okawa/battleship/internal/entity"
	"errors"
)

// TODO persist it somewhere
type GameState struct {
	Id       string       `json:"id"`
	Board    entity.Board `json:"board"`
	Finished bool         `json:"finished"`
}

func (gs *GameState) GetId() string {
	return gs.Id
}

func (gs *GameState) SetId(id string) {
	gs.Id = id
}

type Game struct {
	Persistence GameStatePersistence
}

func (g *Game) Start() (*GameState, error) {
	state := GameState{}
	state.Board = entity.Init()
	state.Finished = false

	if err := g.Persistence.SaveGameState(&state); err != nil {
		return nil, err
	}
	return &state, nil
}

// Receives game id and row/col to shoot
// Returns hit, remaining ships, board and error if happened
func (g *Game) Shoot(id string, row, col int) (bool, int, *GameState, error) {
	state, err := g.Persistence.GetGameState(id)
	if err != nil {
		return false, 0, nil, err
	}

	if state.Finished {
		return false, 0, nil, errors.New("Game finished")
	}

	hit, ships := state.Board.Shoot(row, col)
	if ships == 0 {
		state.Finished = true
	}

	err = g.Persistence.SaveGameState(state)
	if err != nil {
		return false, 0, nil, err
	}

	return hit, ships, state, nil
}
