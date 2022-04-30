package gamestate

import "time"

type Result int

type History struct {
	Id           string
	PlayerTurnId string
	Row          int
	Col          int
	Hit          bool
	Time         time.Time
}

func NewHistory(id, turnId string, row, col int, hit bool, time time.Time) History {
	return History{
		Id:           id,
		PlayerTurnId: turnId,
		Row:          row,
		Col:          col,
		Hit:          hit,
		Time:         time,
	}
}
