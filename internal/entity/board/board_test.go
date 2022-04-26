package board

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew_Size(t *testing.T) {
	assert := assert.New(t)
	board := New(8, 3)

	placementCorrect := true

	size := board.Size
	placementCorrect = placementCorrect && len(board.Placement) == size
	for _, row := range board.Placement {
		placementCorrect = placementCorrect && len(row) == size
	}
	assert.Truef(placementCorrect, "Placement does not have correct square size of '%d'", size)

	stateCorrect := true
	stateCorrect = stateCorrect && len(board.State) == size
	for _, row := range board.State {
		stateCorrect = stateCorrect && len(row) == size
	}
	assert.Truef(stateCorrect, "State does not have correct square size of '%d'", size)
}

func TestNew_ShipCount(t *testing.T) {
	assert := assert.New(t)
	board := New(8, 3)

	count := 0

	for row := 0; row < board.Size; row++ {
		for col := 0; col < board.Size; col++ {
			if board.Placement[row][col] == SHIP {
				count++
			}
		}
	}
	assert.Equal(board.ShipCount, count)
}

func TestShoot_Miss(t *testing.T) {
	assert := assert.New(t)
	board := New(8, 3)

	initialShips := board.ShipCount
	missRow := -1
	missCol := -1

	for row := 0; missRow == -1 && row < board.Size; row++ {
		for col := 0; missCol == -1 && col < board.Size; col++ {
			if board.Placement[row][col] != SHIP {
				missRow = row
				missCol = col
			}
		}
	}

	hit, shipCount := board.Shoot(missRow, missCol)
	assert.False(hit, "It should have missed the shot")
	assert.Equal(initialShips, shipCount, "Ship count should not have changed after a miss shot")
	assert.Equal(MISS, board.State[missRow][missCol], "It should have updated 'state' property with miss")

	for row := 0; row < board.Size; row++ {
		for col := 0; col < board.Size; col++ {
			if row != missRow && col != missCol {
				assert.NotEqual(
					MISS,
					board.State[row][col],
					"It should not have change 'state' property of other placements",
				)
			}
		}
	}
}

func TestShoot_Hit(t *testing.T) {
	assert := assert.New(t)
	board := New(8, 3)

	initialShips := board.ShipCount
	hitRow := -1
	hitCol := -1

	for row := 0; hitRow == -1 && row < board.Size; row++ {
		for col := 0; hitCol == -1 && col < board.Size; col++ {
			if board.Placement[row][col] == SHIP {
				hitRow = row
				hitCol = col
			}
		}
	}

	hit, shipCount := board.Shoot(hitRow, hitCol)
	assert.True(hit, "It should have hit the shot")
	assert.Equal(initialShips, shipCount+1, "Ship count should have been reduced by 1")
	assert.Equal(HIT, board.State[hitRow][hitCol], "It should have updated .state with hit")

	for row := 0; row < board.Size; row++ {
		for col := 0; col < board.Size; col++ {
			if row != hitRow && col != hitCol {
				assert.NotEqual(
					HIT,
					board.State[row][col],
					"It should not have change 'state' property of other placements",
				)
			}
		}
	}
}
