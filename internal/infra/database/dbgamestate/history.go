package dbgamestate

import "elton-okawa/battleship/internal/entity/gamestate"

type History struct {
	Id           string `json:"id"`
	GameStateId  string `json:"gameStateId"`
	PlayerTurnId string `json:"playerTurnId"`
	Row          int    `json:"row"`
	Col          int    `json:"col"`
	Hit          bool   `json:"hit"`
}

func (h *History) GetId() string {
	return h.Id
}

func (h *History) SetId(id string) {
	h.Id = id
}

func historyEntityToDb(gsId string, h []gamestate.History) []History {
	hist := make([]History, len(h))

	for i, ent := range h {
		hist[i] = History{
			Id:           ent.Id,
			GameStateId:  gsId,
			PlayerTurnId: ent.PlayerTurnId,
			Row:          ent.Row,
			Col:          ent.Col,
			Hit:          ent.Hit,
		}
	}

	return hist
}
