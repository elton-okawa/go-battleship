package board

import (
	"math/rand"
)

type Board struct {
	Placement [][]uint8
	State     [][]uint8
	Size      int
	ShipCount int
}

const (
	EMPTY uint8 = 0
	HIT   uint8 = 1
	MISS  uint8 = 2
	SHIP  uint8 = 3
)

func New(size, shipCount int) *Board {
	placement := emptyMap(size)
	state := emptyMap(size)

	board := Board{
		Placement: placement,
		State:     state,
		Size:      size,
		ShipCount: shipCount,
	}

	for i := 0; i < board.ShipCount; i++ {
		placeShip(&board, SHIP)
	}

	return &board
}

func emptyMap(size int) [][]uint8 {
	m := make([][]uint8, size)

	for i := 0; i < size; i++ {
		m[i] = make([]uint8, size)
	}

	// Coincidently the EMPTY is the same as the zeroed value :)
	return m
}

func (board *Board) Shoot(row, col int) (bool, int) {
	hit := board.Placement[row][col] == SHIP
	if hit {
		board.State[row][col] = HIT

		// TODO bigger ships are not sinked right away
		board.ShipCount -= 1
	} else {
		board.State[row][col] = MISS
	}

	return hit, board.ShipCount
}

func placeShip(board *Board, ship uint8) {
	// for now simple one square ships

	positioned := false

	for !positioned {
		row := rand.Intn(board.Size)
		col := rand.Intn(board.Size)

		if board.Placement[row][col] == EMPTY {
			board.Placement[row][col] = ship
			positioned = true
		}
	}
}
