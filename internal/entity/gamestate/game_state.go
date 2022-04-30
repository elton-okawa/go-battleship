package gamestate

import (
	"elton-okawa/battleship/internal/entity/player"
)

type GameState struct {
	Id           string
	PlayerOne    player.Player
	PlayerTwo    player.Player
	History      []History
	PlayerTurnId string
	Finished     bool
}

func New(id string, one, two player.Player, history []History, turn string, finished bool) GameState {
	return GameState{
		Id:           id,
		PlayerOne:    one,
		PlayerTwo:    two,
		History:      history,
		PlayerTurnId: turn,
		Finished:     finished,
	}
}
