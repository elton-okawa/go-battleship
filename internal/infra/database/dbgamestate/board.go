package dbgamestate

import (
	"elton-okawa/battleship/internal/entity/board"
)

type Board struct {
	Id              string `json:"id"`
	Size            int    `json:"size"`
	ShipCoordinates []int  `json:"shipCoordinates"`
}

func (b *Board) GetId() string {
	return b.Id
}

func (b *Board) SetId(id string) {
	b.Id = id
}

func boardEntityToDb(b *board.Board) Board {
	return Board{
		Id:              b.Id,
		Size:            b.Size,
		ShipCoordinates: shipCoordinates(b.InitialShipCount, b.Placement),
	}
}

func shipCoordinates(ships int, p [][]uint8) []int {
	coord := make([]int, ships*2)
	index := 0

	for row := 0; row < len(p); row++ {
		for col := 0; col < len(p[row]); col++ {
			if p[row][col] == board.SHIP {
				coord[index] = row
				coord[index+1] = col
				index += 2
			}
		}
	}

	return coord
}

func (b *Board) ToEntity(h []Turn) *board.Board {
	placement := recreatePlacement(b.Size, b.ShipCoordinates)
	initialShips := len(b.ShipCoordinates) / 2
	state, shipCount := recreateState(placement, h, initialShips)

	return board.New(
		b.Id,
		placement,
		state,
		b.Size,
		initialShips,
		shipCount,
	)
}

func recreatePlacement(size int, shipCoord []int) [][]uint8 {
	placement := make([][]uint8, size)
	for i := 0; i < size; i++ {
		placement[i] = make([]uint8, size)
	}

	for i := 0; i < len(shipCoord)-1; i += 2 {
		placement[shipCoord[i]][shipCoord[i+1]] = board.SHIP
	}

	return placement
}

func recreateState(placement [][]uint8, hist []Turn, initialShips int) (state [][]uint8, shipCount int) {
	state = make([][]uint8, len(placement))
	shipCount = initialShips

	for i := 0; i < len(placement); i++ {
		state[i] = make([]uint8, len(placement[i]))
	}

	for _, t := range hist {
		mark := board.MISS
		if t.Hit {
			mark = board.HIT
			shipCount -= 1
		}

		state[t.Row][t.Col] = mark
	}

	return state, shipCount
}
