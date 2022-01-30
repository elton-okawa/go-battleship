package engine

import (
	"errors"
	"fmt"
)

// TODO persist it somewhere
type GameState struct {
	board    Board
	finished bool
}

var state GameState = GameState{finished: true}

func StartGame() {
	if !state.finished {
		fmt.Println("Game in progress")
	}

	state.board = Init()
	state.finished = false
	fmt.Println(state.board)
}

func Shoot(row, col int) (bool, int, *Board, error) {
	if state.finished {
		return false, 0, nil, errors.New("game finished or not started")
	}

	hit, ships := state.board.Shoot(row, col)
	if ships == 0 {
		state.finished = true
	}
	return hit, ships, &state.board, nil
}
