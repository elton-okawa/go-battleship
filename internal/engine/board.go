package engine

import (
	"math/rand"
	"strings"
)

type Board struct {
	placement [][]rune
	state     [][]rune
	size      int
	shipCount int
}

var EMPTY = '-'
var SINGLE_SQUARE_SHIP = 'S'
var HIT = 'X'
var MISS = '0'

func Init() Board {
	size := 8
	shipCount := 3
	placement := emptyMap(size)
	state := emptyMap(size)

	board := Board{
		placement: placement,
		state:     state,
		size:      size,
		shipCount: shipCount,
	}

	for i := 0; i < board.shipCount; i++ {
		placeShip(&board, SINGLE_SQUARE_SHIP, 1)
	}

	return board
}

func emptyMap(size int) [][]rune {
	m := make([][]rune, size)

	for i := 0; i < size; i++ {
		m[i] = make([]rune, size)
	}

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			m[i][j] = EMPTY
		}
	}

	return m
}

func (board *Board) Shoot(row, col int) (bool, int) {
	hit := board.placement[row][col] == SINGLE_SQUARE_SHIP
	if hit {
		board.state[row][col] = HIT

		// TODO bigger ships are not sinked right away
		board.shipCount -= 1
	} else {
		board.state[row][col] = MISS
	}

	return hit, board.shipCount
}

func placeShip(board *Board, char rune, size int) {
	// for now simple one square ships

	positioned := false

	for !positioned {
		row := rand.Intn(board.size)
		col := rand.Intn(board.size)

		if board.placement[row][col] == EMPTY {
			board.placement[row][col] = char
			positioned = true
		}
	}
}

func (b Board) String() string {
	var sb strings.Builder
	for i := 0; i < b.size; i++ {
		for j := 0; j < b.size; j++ {
			sb.WriteRune(b.placement[i][j])
			sb.WriteRune(' ')
		}
		sb.WriteRune('\n')
	}

	return sb.String()
}
