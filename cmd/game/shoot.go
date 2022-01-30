package main

import (
	"elton-okawa/battleship/internal/engine"
	"errors"
	"fmt"
	"strconv"
)

type Shoot struct {
	row int
	col int
}

func (s *Shoot) Parse(args []string) error {
	if len(args) != 2 {
		return errors.New("shoot command must receive exactly 2 arguments")
	}

	row, err := strconv.Atoi(args[0])
	if err != nil {
		return errors.New("row must be a valid integer")
	}

	col, err := strconv.Atoi(args[1])
	if err != nil {
		return errors.New("col must be a valid integer")
	}

	s.row = row
	s.col = col

	return nil
}

func (s *Shoot) Execute() (bool, error) {
	hit, ships, board, err := engine.Shoot(s.row, s.col)
	if err != nil {
		return false, err
	} else {
		fmt.Println(board)
		fmt.Printf("Your shot hit: %t\n", hit)
		fmt.Printf("There is/are %d ships squares remaining\n", ships)

		if ships == 0 {
			fmt.Printf("Game finished, you can start another one\n")
		}

		return ships == 0, err
	}
}
