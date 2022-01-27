package main

type Command interface {
	Parse([]string) error
	Execute()
}

var Commands = map[string]Command{
	"start": Start{},
}
