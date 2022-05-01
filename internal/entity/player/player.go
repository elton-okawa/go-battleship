package player

import (
	"elton-okawa/battleship/internal/entity/board"
)

type Player struct {
	Id    string
	Name  string
	Board *board.Board
}

func New(id, name string, b *board.Board) Player {
	return Player{
		Id:    id,
		Name:  name,
		Board: b,
	}
}
