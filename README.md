# Go Battleship game

Implementation of [florinpop - App Ideas - Battleship Game Engine](https://github.com/florinpop17/app-ideas/blob/master/Projects/3-Advanced/Battleship-Game-Engine.md)

## How to play

Run
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
go test ./**/*
```

## TODO

* Random placement bigger ships - now we only place single square
* Keep track if an entire ship was sinked (more than one square)