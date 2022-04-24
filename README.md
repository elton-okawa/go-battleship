# Go Battleship game

This repository has the intention of studying Go Programming Language and [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)

Implementation of [florinpop - App Ideas - Battleship Game Engine](https://github.com/florinpop17/app-ideas/blob/master/Projects/3-Advanced/Battleship-Game-Engine.md)

Take a look at:
- [Projects Tab](https://github.com/elton-okawa/go-battleship/projects/1) to view the planned backlog
- [Architecture and control flow](./docs/architecture.md)
- [About current third party packages in use](./docs/third-party-packages.md)

## Start server

Run server
```
go run ./cmd/server/main.go
```

## Testing

```
go test ./... -v
```

### Coverage
```
go test ./... -coverprofile cover.out
go tool cover -html=cover.out -o cover.html
```

## Generate OpenApi code
```
go generate ./...
```
