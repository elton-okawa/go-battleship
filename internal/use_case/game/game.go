package game

import (
	"elton-okawa/battleship/internal/entity"
	use_case_errors "elton-okawa/battleship/internal/use_case/errors"
	"errors"
	"fmt"
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

func NewGameUseCase(gsp GameStatePersistence) GameUseCase {
	return GameUseCase{
		persistence: gsp,
	}
}

type GameUseCase struct {
	persistence GameStatePersistence
}

func (g GameUseCase) Start(gob GameOutputBoundary) {
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
func (g GameUseCase) Shoot(gob GameOutputBoundary, id string, row, col int) {
	state, err := g.persistence.GetGameState(id)
	if err != nil {
		notFoundErr := use_case_errors.NewError(
			fmt.Sprintf("Could not find game with id '%s'\n%v", id, err),
			use_case_errors.ElementNotFound,
			nil,
		)
		gob.ShootResult(nil, false, 0, notFoundErr)
		return
	}

	if state.Finished {
		gob.ShootResult(nil, false, 0, errors.New("game finished"))
		return
	}

	// TODO check if row and col are valid (not shot or inside board boundaries)
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
