package gamestate

import (
	"elton-okawa/battleship/internal/entity/board"
)

type GameState struct {
	Id           string
	AccountOneId string
	AccountTwoId string
	BoardOne     *board.Board
	BoardTwo     *board.Board
	History      []Turn
	PlayerTurnId string
	Finished     bool
}

func New(id, accOneId, accTwoId string, bOne, bTwo *board.Board, history []Turn, turn string, finished bool) GameState {
	return GameState{
		Id:           id,
		AccountOneId: accOneId,
		AccountTwoId: accTwoId,
		BoardOne:     bOne,
		BoardTwo:     bTwo,
		History:      history,
		PlayerTurnId: turn,
		Finished:     finished,
	}
}
