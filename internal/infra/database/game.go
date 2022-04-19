package database

import "elton-okawa/battleship/internal/usecase/game"

type GameDao struct {
	driver JsonDatabase
}

func NewGameDao(path string) GameDao {
	return GameDao{
		driver: NewJsonDatabase(path),
	}
}

func (db GameDao) SaveGameState(gs *game.GameState) error {
	return db.driver.Save(gs)
}

func (db GameDao) GetGameState(id string) (*game.GameState, error) {
	gs := game.GameState{}
	err := db.driver.Get(id, &gs)

	return &gs, err
}
