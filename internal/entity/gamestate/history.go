package gamestate

type Result int

type History []Turn

type Turn struct {
	Id           string
	PlayerTurnId string
	Row          int
	Col          int
	Hit          bool
	Time         int64
}

func NewTurn(id, turnId string, row, col int, hit bool, time int64) Turn {
	return Turn{
		Id:           id,
		PlayerTurnId: turnId,
		Row:          row,
		Col:          col,
		Hit:          hit,
		Time:         time,
	}
}
