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
