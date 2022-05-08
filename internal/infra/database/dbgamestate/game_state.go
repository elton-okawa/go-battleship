package dbgamestate

import (
	"elton-okawa/battleship/internal/entity/gamestate"
	"elton-okawa/battleship/internal/infra/database"
	"fmt"
	"sort"
)

type GameState struct {
	Id           string `json:"id"`
	AccountOneId string `json:"accountOneId"`
	AccountTwoId string `json:"accountTwoId"`
	BoardOneId   string `json:"boardOneId"`
	BoardTwoId   string `json:"boardTwoId"`
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

func (db Repository) Get(id string) (gamestate.GameState, error) {
	var gsData GameState
	if err := db.driver.Get(id, &gsData); err != nil {
		return gamestate.GameState{}, fmt.Errorf("error while reading game state: %w", err)
	}

	var bOneData Board
	if err := db.driver.Get(gsData.BoardOneId, &bOneData); err != nil {
		return gamestate.GameState{}, fmt.Errorf("error while reading board one: %w", err)
	}

	var bTwoData Board
	if err := db.driver.Get(gsData.BoardTwoId, &bTwoData); err != nil {
		return gamestate.GameState{}, fmt.Errorf("error while reading board two: %w", err)
	}

	var history History
	if err := db.driver.FindAllBy("gameStateId", gsData.Id, &history); err != nil {
		return gamestate.GameState{}, fmt.Errorf("error while reading history: %w", err)
	}
	sort.Sort(history)

	targetOne, targetTwo := splitHistory(gsData.AccountOneId, gsData.AccountTwoId, history)

	gs := gamestate.New(
		gsData.Id,
		gsData.AccountOneId,
		gsData.AccountTwoId,
		bOneData.ToEntity(targetOne),
		bTwoData.ToEntity(targetTwo),
		mapHistory(history),
		gsData.PlayerTurnId,
		gsData.Finished,
	)

	return gs, nil
}

func (db Repository) Save(gs gamestate.GameState) error {
	data := GameState{
		Id:           gs.Id,
		AccountOneId: gs.AccountOneId,
		AccountTwoId: gs.AccountTwoId,
		BoardOneId:   gs.BoardOne.Id,
		BoardTwoId:   gs.BoardTwo.Id,
		PlayerTurnId: gs.PlayerTurnId,
		Finished:     gs.Finished,
	}

	bOne := boardEntityToDb(gs.BoardOne)
	bTwo := boardEntityToDb(gs.BoardTwo)

	hist := historyEntityToDb(gs.Id, gs.History)

	entities := []database.Entity{&data, &bOne, &bTwo}
	for i := range hist {
		entities = append(entities, &hist[i])
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

func splitHistory(pOne, pTwo string, hist []Turn) (histOne, histTwo []Turn) {
	for _, h := range hist {
		if h.PlayerTurnId == pOne {
			histTwo = append(histTwo, h)
		} else {
			histOne = append(histOne, h)
		}
	}

	return histOne, histTwo
}

func mapHistory(hist History) gamestate.History {
	entities := make([]gamestate.Turn, len(hist))
	for i, h := range hist {
		entities[i] = h.ToEntity()
	}

	return entities
}
