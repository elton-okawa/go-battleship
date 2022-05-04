package main

import (
	"elton-okawa/battleship/internal/infra/database/dbgamestate"
	"elton-okawa/battleship/internal/interface_adapter/presenter"
)

type Command interface {
	Parse([]string) error
	Execute() (bool, error)
	Description() string
}

var cmdPresenter = &presenter.CommandLinePresenter{}
var cmdPersistence = &dbgamestate.Repository{}

var Commands = map[string]Command{
	// "start": &Start{
	// 	persistence: cmdPersistence,
	// 	presenter:   cmdPresenter,
	// },
	// "shoot": &Shoot{
	// 	persistence: cmdPersistence,
	// 	presenter:   cmdPresenter,
	// },
}
