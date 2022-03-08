package database

import "elton-okawa/battleship/internal/use_case"

type Database struct {
	driver JsonDatabase
}

// TODO who should instantiate it?
var DefaultDatabase Database = Database{driver: JsonDatabase{Filepath: "./db/games.json"}}

func (db *Database) SaveGameState(gs *use_case.GameState) error {
	return db.driver.Save(gs)
}

func (db *Database) GetGameState(id string) (*use_case.GameState, error) {
	gs := use_case.GameState{}
	err := db.driver.Get(id, &gs)

	return &gs, err
}
