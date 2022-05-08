package board

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew_Size(t *testing.T) {
	assert := assert.New(t)
	board := Initialize(8, 3)

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
	board := Initialize(8, 3)

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
	board := Initialize(8, 3)
	initialShips := board.ShipCount

	missRow, missCol := find(board.Placement, EMPTY)
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
	board := Initialize(8, 3)
	initialShips := board.ShipCount

	hitRow, hitCol := find(board.Placement, SHIP)
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

func TestCanShoot(t *testing.T) {
	assert := assert.New(t)

	board := Initialize(8, 3)
	board.State[1][0] = HIT
	board.State[1][1] = MISS

	emptyRow, emptyCol := find(board.State, EMPTY)
	assert.True(board.CanShoot(emptyRow, emptyCol), "It should be possible to shot a empty spot")

	hitRow, hitCol := find(board.State, HIT)
	assert.False(board.CanShoot(hitRow, hitCol), "It should not be possible to shot in a hit spot")

	missRow, missCol := find(board.State, MISS)
	assert.False(board.CanShoot(missRow, missCol), "It should not be possible to shot in a miss spot")
}

func find(board [][]uint8, target uint8) (int, int) {
	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board[row]); col++ {
			if board[row][col] == target {
				return row, col
			}
		}
	}

	// in test environment it should not happen
	return -1, -1
}
