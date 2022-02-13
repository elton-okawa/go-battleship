# Go Battleship game

Implementation of [florinpop - App Ideas - Battleship Game Engine](https://github.com/florinpop17/app-ideas/blob/master/Projects/3-Advanced/Battleship-Game-Engine.md)

## How to play

Run server
```
go run ./cmd/server
```

Run CLI
```
go run ./cmd/game/*
```


## Rules

Ships
* Destroyer - 2 squares
* Cruiser - 3 squares
* Battleship - 4 squares

## Testing

```
go test ./**/* -v
```

### Coverage
```
go test ./**/* -coverprofile cover.out
go tool cover -html=cover.out -o cover.html
```

## TODO

* Create account
* Persist each game state in a DS
* Persist player statistic over games
* Random placement bigger ships - now we only place single square
* Keep track if an entire ship was sinked (more than one square)