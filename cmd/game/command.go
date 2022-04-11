package main

import (
	"elton-okawa/battleship/internal/database"
	"elton-okawa/battleship/internal/interface_adapter/presenter"
)

type Command interface {
	Parse([]string) error
	Execute() (bool, error)
	Description() string
}

var cmdPresenter = &presenter.CommandLinePresenter{}
var cmdPersistence = &database.GameDao{}

var Commands = map[string]Command{
	"start": &Start{
		persistence: cmdPersistence,
		presenter:   cmdPresenter,
	},
	"shoot": &Shoot{
		persistence: cmdPersistence,
		presenter:   cmdPresenter,
	},
}
