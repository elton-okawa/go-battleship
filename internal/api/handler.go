package api

type Handler interface {
	Parse(string) error
	Execute()
}
