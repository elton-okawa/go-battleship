package main

import (
	"elton-okawa/battleship/internal/use_case"
	"errors"
	"strconv"
)

type Shoot struct {
	persistence use_case.GameStatePersistence
	presenter   use_case.GameOutputBoundary
	row         int
	col         int
	id          string
}

func (s *Shoot) Description() string {
	return "<row> <col> - shoot at <row> <col>"
}

func (s *Shoot) Parse(args []string) error {
	if len(args) != 3 {
		return errors.New("shoot command must receive exactly 3 arguments")
	}

	id := args[0]

	row, err := strconv.Atoi(args[1])
	if err != nil {
		return errors.New("row must be a valid integer")
	}

	col, err := strconv.Atoi(args[2])
	if err != nil {
		return errors.New("col must be a valid integer")
	}

	s.id = id
	s.row = row
	s.col = col

	return nil
}

func (s *Shoot) Execute() (bool, error) {
	game := use_case.NewGame(s.persistence)
	game.Shoot(s.presenter, s.id, s.row, s.col)

	return true, nil
	// if err != nil {
	// 	return false, err
	// } else {
	// 	fmt.Println(board)
	// 	fmt.Printf("Your shot hit: %t\n", hit)
	// 	fmt.Printf("There is/are %d ships squares remaining\n", ships)

	// 	if ships == 0 {
	// 		fmt.Printf("Game finished, you can start another one\n")
	// 	}

	// 	return ships == 0, err
	// }
}
