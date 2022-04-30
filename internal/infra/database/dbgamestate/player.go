package dbgamestate

import (
	"elton-okawa/battleship/internal/entity/board"
	"elton-okawa/battleship/internal/entity/gamestate"
	"elton-okawa/battleship/internal/entity/player"
)

type Player struct {
	Id        string `json:"id"`
	AccountId string `json:"accountId"`
	BoardId   string `json:"boardId"`
}

func (p *Player) GetId() string {
	return p.Id
}

func (p *Player) SetId(id string) {
	p.Id = id
}

func playerEntityToDb(p player.Player) Player {
	return Player{
		Id:        p.Id,
		AccountId: p.Account.Id,
		BoardId:   p.Board.Id,
	}
}

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
