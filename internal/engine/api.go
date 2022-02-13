package engine

import (
	"errors"
	"fmt"
)

// TODO persist it somewhere
type GameState struct {
	Board    Board
	Finished bool
}

var state GameState = GameState{Finished: true}

func StartGame() GameState {
	if !state.Finished {
		fmt.Println("Game in progress")
	}

	state.Board = Init()
	state.Finished = false
	fmt.Println(state.Board)

	return state
}

func Shoot(row, col int) (bool, int, *Board, error) {
	if state.Finished {
		return false, 0, nil, errors.New("game finished or not started")
	}

	hit, ships := state.Board.Shoot(row, col)
	if ships == 0 {
		state.Finished = true
	}
	return hit, ships, &state.Board, nil
}
