package dbgamestate

import (
	"elton-okawa/battleship/internal/entity/gamestate"
	"elton-okawa/battleship/internal/infra/database"
)

type GameState struct {
	Id           string `json:"id"`
	PlayerOneId  string `json:"playerOneId"`
	PlayerTwoId  string `json:"playerTwoId"`
	PlayerTurnId string `json:"turnId"`
	Finished     bool   `json:"finished"`
}

func (gs *GameState) GetId() string {
	return gs.Id
}

func (gs *GameState) SetId(id string) {
	gs.Id = id
}

type Repository struct {
	driver database.JsonDatabase
}

func New(filepath string) Repository {
	return Repository{
		driver: database.NewJsonDatabase(filepath),
	}
}

func (db Repository) Get(id string) (*gamestate.GameState, error) {
	var data GameState
	err := db.driver.Get(id, &data)
	if err != nil {
		return nil, err
	}

	// gs := gamestate.New(
	// 	data.Id,

	// 	data.PlayerTurnId,
	// 	data.Finished,
	// )

	return nil, nil
}

func (db Repository) Save(gs gamestate.GameState) error {
	data := GameState{
		Id:           gs.Id,
		PlayerOneId:  gs.PlayerOne.Id,
		PlayerTwoId:  gs.PlayerTwo.Id,
		PlayerTurnId: gs.PlayerTurnId,
		Finished:     gs.Finished,
	}

	pOne := playerEntityToDb(gs.PlayerOne)
	bOne := boardEntityToDb(gs.PlayerOne.Board)

	pTwo := playerEntityToDb(gs.PlayerTwo)
	bTwo := boardEntityToDb(gs.PlayerTwo.Board)

	hist := historyEntityToDb(gs.Id, gs.History)

	entities := []database.Entity{&data, &pOne, &bOne, &pTwo, &bTwo}
	for _, h := range hist {
		entities = append(entities, &h)
	}

	// TODO transaction and distinct kind
	return saveMultiple(db.driver, entities)
}

func saveMultiple(driver database.JsonDatabase, data []database.Entity) error {
	for _, d := range data {
		if err := driver.Save(d); err != nil {
			return err
		}
	}

	return nil
}
