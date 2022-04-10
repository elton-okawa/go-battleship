package entity

import (
	"testing"
)

func TestNewBoardSize(t *testing.T) {
	board := NewBoard(8, 3)

	placementCorrect := true

	size := board.Size
	placementCorrect = placementCorrect && len(board.Placement) == size
	for _, row := range board.Placement {
		placementCorrect = placementCorrect && len(row) == size
	}

	if !placementCorrect {
		t.Errorf("Placement does not have correct square size of '%d'", size)
	}

	stateCorrect := true
	stateCorrect = stateCorrect && len(board.State) == size
	for _, row := range board.State {
		stateCorrect = stateCorrect && len(row) == size
	}

	if !stateCorrect {
		t.Errorf("State does not have correct square size of '%d'", size)
	}
}

func TestNewBoardShipCount(t *testing.T) {
	board := NewBoard(8, 3)

	count := 0

	for row := 0; row < board.Size; row++ {
		for col := 0; col < board.Size; col++ {
			if board.Placement[row][col] == SINGLE_SQUARE_SHIP {
				count++
			}
		}
	}

	if count != board.ShipCount {
		t.Errorf("Ship count does not match %d", board.ShipCount)
	}
}

func TestShootMiss(t *testing.T) {
	board := NewBoard(8, 3)

	initialShips := board.ShipCount
	missRow := -1
	missCol := -1

	for row := 0; missRow == -1 && row < board.Size; row++ {
		for col := 0; missCol == -1 && col < board.Size; col++ {
			if board.Placement[row][col] != SINGLE_SQUARE_SHIP {
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

	if board.State[missRow][missCol] != MISS {
		t.Errorf("It should have updated .state with miss")
	}

	for row := 0; row < board.Size; row++ {
		for col := 0; col < board.Size; col++ {
			if row != missRow && col != missCol && board.State[row][col] == MISS {
				t.Errorf("It should not have change .state of other placements")
			}
		}
	}
}

func TestShootHit(t *testing.T) {
	board := NewBoard(8, 3)

	initialShips := board.ShipCount
	hitRow := -1
	hitCol := -1

	for row := 0; hitRow == -1 && row < board.Size; row++ {
		for col := 0; hitCol == -1 && col < board.Size; col++ {
			if board.Placement[row][col] == SINGLE_SQUARE_SHIP {
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

	if board.State[hitRow][hitCol] != HIT {
		t.Errorf("It should have updated .state with hit")
	}

	for row := 0; row < board.Size; row++ {
		for col := 0; col < board.Size; col++ {
			if row != hitRow && col != hitCol && board.State[row][col] == HIT {
				t.Errorf("It should not have change .state of other placements")
			}
		}
	}
}
