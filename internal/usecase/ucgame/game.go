package ucgame

import (
	"elton-okawa/battleship/internal/entity/board"
	"elton-okawa/battleship/internal/entity/gamerequest"
	"elton-okawa/battleship/internal/entity/gamestate"
	"elton-okawa/battleship/internal/usecase/ucerror"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

func New(gsRepo GameStateRepository, grRepo GameRequestRepository) UseCase {
	return UseCase{
		gsRepo: gsRepo,
		grRepo: grRepo,
	}
}

type UseCase struct {
	gsRepo GameStateRepository
	grRepo GameRequestRepository
}

type GameRequestRepository interface {
	FindOwn(owner string) (*gamerequest.GameRequest, error)
	FindPending() (*gamerequest.GameRequest, error)
	Save(gs *gamerequest.GameRequest) error
}

func (uc UseCase) Start(gob GameOutputBoundary, pId string) {
	// TODO do not create a new game if player already have a request
	// ownGr, err := uc.grRepo.FindOwn(pId)
	// if err != nil {
	// 	// handle error
	// }

	// TODO matchmaking
	gr, err := uc.grRepo.FindPending()
	if err != nil {
		fmt.Printf("error finding gs: %v\n", err)
		// handle error
	}

	if gr != nil {
		gr.ChallengerId = pId
		gr.Pending = false

		// TODO transaction?
		uc.grRepo.Save(gr)

		gs := gamestate.New(
			uuid.NewString(),
			gr.OwnerId,
			gr.ChallengerId,
			board.New(8, 3),
			board.New(8, 3),
			[]gamestate.History{},
			gr.OwnerId,
			false,
		)

		// handle error
		uc.gsRepo.Save(gs)
	} else {
		gr := gamerequest.New(uuid.NewString(), pId, "", true)
		uc.grRepo.Save(&gr)
	}

	// if err := g.persistence.SaveGameState(&state); err != nil {
	// 	gob.StartResult(nil, err)
	// } else {
	// 	gob.StartResult(&state, nil)
	// }
}

// Receives game id and row/col to shoot
func (g UseCase) Shoot(gob GameOutputBoundary, id string, row, col int) {
	state, err := g.gsRepo.Get(id)
	if err != nil {
		notFoundErr := ucerror.New(
			fmt.Sprintf("Could not find game with id '%s'\n%v", id, err),
			ucerror.ElementNotFound,
			nil,
		)
		gob.ShootResult(nil, false, 0, notFoundErr)
		return
	}

	if state.Finished {
		gob.ShootResult(nil, false, 0, errors.New("game finished"))
		return
	}

	// TODO check if row and col are valid (not shot or inside board boundaries)
	// hit, ships := state.Board.Shoot(row, col)
	// if ships == 0 {
	// 	state.Finished = true
	// }

	// err = g.persistence.SaveGameState(state)
	// if err != nil {
	// 	gob.ShootResult(nil, false, 0, err)
	// 	return
	// }

	// gob.ShootResult(state, hit, ships, nil)
}
