package use_case

import (
	"elton-okawa/battleship/internal/entity"
	"errors"
)

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

func NewGame(gsp GameStatePersistence) *Game {
	return &Game{
		persistence: gsp,
	}
}

type Game struct {
	persistence GameStatePersistence
}

func (g *Game) Start(gob GameOutputBoundary) {
	state := GameState{}
	state.Board = entity.Init()
	state.Finished = false

	if err := g.persistence.SaveGameState(&state); err != nil {
		gob.StartResult(nil, err)
	} else {
		gob.StartResult(&state, nil)
	}
}

// Receives game id and row/col to shoot
func (g *Game) Shoot(gob GameOutputBoundary, id string, row, col int) {
	state, err := g.persistence.GetGameState(id)
	if err != nil {
		gob.ShootResult(nil, false, 0, err)
		return
	}

	if state.Finished {
		gob.ShootResult(nil, false, 0, errors.New("Game finished"))
		return
	}

	hit, ships := state.Board.Shoot(row, col)
	if ships == 0 {
		state.Finished = true
	}

	err = g.persistence.SaveGameState(state)
	if err != nil {
		gob.ShootResult(nil, false, 0, err)
		return
	}

	gob.ShootResult(state, hit, ships, nil)
}
