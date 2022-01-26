package engine

import (
	"math/rand"
	"strings"
)

type Board struct {
	board [][]rune
	size  int
}

var EMPTY = '-'

func Init() Board {
	size := 8
	b := make([][]rune, size)

	for i := 0; i < size; i++ {
		b[i] = make([]rune, size)
	}

	board := Board{
		board: b,
		size:  size,
	}

	for i := 0; i < board.size; i++ {
		for j := 0; j < board.size; j++ {
			board.board[i][j] = EMPTY
		}
	}

	for i := 0; i < 3; i++ {
		placeShip(&board, 'S', 1)
	}

	return board
}

func placeShip(board *Board, char rune, size int) {
	// for now simple one square ships

	positioned := false

	for !positioned {
		row := rand.Intn(board.size)
		col := rand.Intn(board.size)

		if board.board[row][col] == EMPTY {
			board.board[row][col] = char
			positioned = true
		}
	}
}

func (b Board) String() string {
	var sb strings.Builder
	for i := 0; i < b.size; i++ {
		for j := 0; j < b.size; j++ {
			sb.WriteRune(b.board[i][j])
			sb.WriteRune(' ')
		}
		sb.WriteRune('\n')
	}

	return sb.String()
}
