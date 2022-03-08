package entity

import (
	"math/rand"
	"strconv"
	"strings"
)

type Board struct {
	Placement [][]rune `json:"placement"`
	State     [][]rune `json:"state"`
	Size      int      `json:"size"`
	ShipCount int      `json:"shipCount"`
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
		Placement: placement,
		State:     state,
		Size:      size,
		ShipCount: shipCount,
	}

	for i := 0; i < board.ShipCount; i++ {
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
	hit := board.Placement[row][col] == SINGLE_SQUARE_SHIP
	if hit {
		board.State[row][col] = HIT

		// TODO bigger ships are not sinked right away
		board.ShipCount -= 1
	} else {
		board.State[row][col] = MISS
	}

	return hit, board.ShipCount
}

func placeShip(board *Board, char rune, size int) {
	// for now simple one square ships

	positioned := false

	for !positioned {
		row := rand.Intn(board.Size)
		col := rand.Intn(board.Size)

		if board.Placement[row][col] == EMPTY {
			board.Placement[row][col] = char
			positioned = true
		}
	}
}

func (b Board) String() string {
	var sb strings.Builder

	// Coordinates helper
	sb.WriteString("\\")
	addNumberRow(&sb, b.Size)
	sb.WriteRune(' ')
	addNumberRow(&sb, b.Size)

	for i := 0; i < b.Size; i++ {
		sb.WriteRune('\n')
		sb.WriteString(strconv.Itoa(i))
		for j := 0; j < b.Size; j++ {
			sb.WriteRune(' ')
			sb.WriteRune(b.Placement[i][j])
		}

		sb.WriteRune(' ')
		for j := 0; j < b.Size; j++ {
			sb.WriteRune(' ')
			sb.WriteRune(b.State[i][j])
		}
	}

	return sb.String()
}

func addNumberRow(sb *strings.Builder, size int) {
	for row := 0; row < size; row++ {
		sb.WriteRune(' ')
		sb.WriteString(strconv.Itoa(row))
	}
}
