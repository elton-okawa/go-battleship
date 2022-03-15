package presenter

import (
	"elton-okawa/battleship/internal/use_case"
	"fmt"
)

type RestApiPresenter struct{}

func NewRestApiPresenter() *RestApiPresenter {
	return &RestApiPresenter{}
}

func (rp *RestApiPresenter) StartResult(gs *use_case.GameState, err error) {
	fmt.Printf("%+v\n", gs)
}

func (rp *RestApiPresenter) ShootResult(gs *use_case.GameState, hit bool, ships int, err error) {
	fmt.Printf("hit: %t, ships: %d\n", hit, ships)
}
