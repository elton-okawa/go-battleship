package dbplayer

import (
	"elton-okawa/battleship/internal/entity/player"
	"elton-okawa/battleship/internal/infra/database"
)

type Player struct {
	Id string `json:"id"`
}

func (p *Player) GetId() string {
	return p.Id
}

func (p *Player) SetId(id string) {
	p.Id = id
}

type Repository struct {
	driver database.JsonDatabase
}

func New(filepath string) Repository {
	return Repository{
		driver: database.NewJsonDatabase(filepath),
	}
}

func (db Repository) Get(id string) (player.Player, error) {
	data, err := db.Get(id)
	if err != nil {
		return player.Player{}, err
	}

	p := player.Player{
		Id: data.Id,
	}

	return p, nil
}
