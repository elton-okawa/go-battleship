package use_case

import (
	"elton-okawa/battleship/internal/entity"
	"errors"
	"fmt"
)

// TODO persist it somewhere
type GameState struct {
	Board    entity.Board
	Finished bool
}

var state GameState = GameState{Finished: true}

func StartGame() GameState {
	if !state.Finished {
		fmt.Println("Game in progress")
	}

	state.Board = entity.Init()
	state.Finished = false
	fmt.Println(state.Board)

	return state
}

func Shoot(row, col int) (bool, int, *entity.Board, error) {
	if state.Finished {
		return false, 0, nil, errors.New("game finished or not started")
	}

	hit, ships := state.Board.Shoot(row, col)
	if ships == 0 {
		state.Finished = true
	}
	return hit, ships, &state.Board, nil
}
