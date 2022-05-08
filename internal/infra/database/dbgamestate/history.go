package dbgamestate

import (
	"elton-okawa/battleship/internal/entity/gamestate"
)

type History []Turn

func (h History) Len() int {
	return len(h)
}

func (h History) Less(i, j int) bool {
	return h[i].Time < h[j].Time
}

func (h History) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

type Turn struct {
	Id           string `json:"id"`
	GameStateId  string `json:"gameStateId"`
	PlayerTurnId string `json:"playerTurnId"`
	Row          int    `json:"row"`
	Col          int    `json:"col"`
	Hit          bool   `json:"hit"`
	Time         int64  `json:"time"`
}

func (h *Turn) GetId() string {
	return h.Id
}

func (h *Turn) SetId(id string) {
	h.Id = id
}

func historyEntityToDb(gsId string, h gamestate.History) History {
	hist := make([]Turn, len(h))

	for i, ent := range h {
		hist[i] = Turn{
			Id:           ent.Id,
			GameStateId:  gsId,
			PlayerTurnId: ent.PlayerTurnId,
			Row:          ent.Row,
			Col:          ent.Col,
			Hit:          ent.Hit,
			Time:         ent.Time,
		}
	}

	return hist
}

func (h Turn) ToEntity() gamestate.Turn {
	return gamestate.NewTurn(
		h.Id,
		h.PlayerTurnId,
		h.Row,
		h.Col,
		h.Hit,
		h.Time,
	)
}
