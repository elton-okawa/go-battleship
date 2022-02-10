package main

type Command interface {
	Parse([]string) error
	Execute() (bool, error)
	Description() string
}

var Commands = map[string]Command{
	"start": &Start{},
	"shoot": &Shoot{},
}
