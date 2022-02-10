package engine

import (
	"testing"
)

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

func TestShootMiss(t *testing.T) {
	board := Init()

	initialShips := board.shipCount
	missRow := -1
	missCol := -1

	for row := 0; missRow == -1 && row < board.size; row++ {
		for col := 0; missCol == -1 && col < board.size; col++ {
			if board.placement[row][col] != SINGLE_SQUARE_SHIP {
				missRow = row
				missCol = col
			}
		}
	}

	hit, shipCount := board.Shoot(missRow, missCol)

	if hit != false {
		t.Errorf("It should have missed the shot")
	}

	if initialShips != shipCount {
		t.Errorf("Ship count should not have changed after a miss shot")
	}

	if board.state[missRow][missCol] != MISS {
		t.Errorf("It should have updated .state with miss")
	}

	for row := 0; row < board.size; row++ {
		for col := 0; col < board.size; col++ {
			if row != missRow && col != missCol && board.state[row][col] == MISS {
				t.Errorf("It should not have change .state of other placements")
			}
		}
	}
}

func TestShootHit(t *testing.T) {
	board := Init()

	initialShips := board.shipCount
	hitRow := -1
	hitCol := -1

	for row := 0; hitRow == -1 && row < board.size; row++ {
		for col := 0; hitCol == -1 && col < board.size; col++ {
			if board.placement[row][col] == SINGLE_SQUARE_SHIP {
				hitRow = row
				hitCol = col
			}
		}
	}

	hit, shipCount := board.Shoot(hitRow, hitCol)

	if hit != true {
		t.Errorf("It should have hit the shot")
	}

	if initialShips != shipCount+1 {
		t.Errorf("Ship count should have been reduced by one (initialShips: %d, currentShips: %d)", initialShips, shipCount)
	}

	if board.state[hitRow][hitCol] != HIT {
		t.Errorf("It should have updated .state with hit")
	}

	for row := 0; row < board.size; row++ {
		for col := 0; col < board.size; col++ {
			if row != hitRow && col != hitCol && board.state[row][col] == HIT {
				t.Errorf("It should not have change .state of other placements")
			}
		}
	}
}
