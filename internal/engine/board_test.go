package engine

import "testing"

func TestInitSize(t *testing.T) {
	board := Init()

	placementCorrect := true

	size := board.size
	placementCorrect = placementCorrect && len(board.placement) == size
	for _, row := range board.placement {
		placementCorrect = placementCorrect && len(row) == size
	}

	if !placementCorrect {
		t.Errorf("Placement does not have correct square size of '%d'", size)
	}

	stateCorrect := true
	stateCorrect = stateCorrect && len(board.state) == size
	for _, row := range board.state {
		stateCorrect = stateCorrect && len(row) == size
	}

	if !stateCorrect {
		t.Errorf("State does not have correct square size of '%d'", size)
	}
}

func TestInitShipCount(t *testing.T) {
	board := Init()

	count := 0

	for row := 0; row < board.size; row++ {
		for col := 0; col < board.size; col++ {
			if board.placement[row][col] == SINGLE_SQUARE_SHIP {
				count++
			}
		}
	}

	if count != board.shipCount {
		t.Errorf("Ship count does not match %d", board.shipCount)
	}
}
