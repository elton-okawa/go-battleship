package dbgamerequest

import (
	"elton-okawa/battleship/internal/entity/gamerequest"
	"elton-okawa/battleship/internal/infra/database"
)

type GameRequest struct {
	Id           string `json:"id"`
	OwnerId      string `json:"ownerId"`
	ChallengerId string `json:"challengerId"`
	Pending      bool   `json:"pending"`
}

func (gr *GameRequest) GetId() string {
	return gr.Id
}

func (gr *GameRequest) SetId(id string) {
	gr.Id = id
}

type Repository struct {
	driver database.JsonDatabase
}

func New(filepath string) Repository {
	return Repository{
		driver: database.NewJsonDatabase(filepath),
	}
}

func (db Repository) FindOwn(owner string) (*gamerequest.GameRequest, error) {
	var data GameRequest
	if err := db.driver.FindFirstBy("ownerId", owner, &data); err != nil {
		return nil, err
	}

	gr := gamerequest.New(data.Id, data.OwnerId, data.ChallengerId, data.Pending)
	return &gr, nil
}

func (db Repository) FindPending() (*gamerequest.GameRequest, error) {
	var data GameRequest
	if err := db.driver.FindFirstBy("pending", true, &data); err != nil {
		return nil, err
	}

	gr := gamerequest.New(data.Id, data.OwnerId, data.ChallengerId, data.Pending)
	return &gr, nil
}

func (db Repository) Save(gr *gamerequest.GameRequest) error {
	data := &GameRequest{
		Id:           gr.Id,
		OwnerId:      gr.OwnerId,
		ChallengerId: gr.ChallengerId,
		Pending:      gr.Pending,
	}

	return db.driver.Save(data)
}
