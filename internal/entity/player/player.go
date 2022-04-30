package player

import (
	"elton-okawa/battleship/internal/entity/account"
	"elton-okawa/battleship/internal/entity/board"
)

type Player struct {
	Id      string
	Account account.Account
	Board   *board.Board
}

func New(id string, acc account.Account, b *board.Board) Player {
	return Player{
		Id:      id,
		Account: acc,
		Board:   b,
	}
}
