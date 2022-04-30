package ucgame

import (
	"elton-okawa/battleship/internal/entity/gamerequest"
	"elton-okawa/battleship/internal/entity/gamestate"
	"elton-okawa/battleship/internal/entity/player"
	"elton-okawa/battleship/internal/usecase/ucerror"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

func New(gsRepo GameStateRepository, grRepo GameRequestRepository, pRepo PlayerRepository) UseCase {
	return UseCase{
		gsRepo: gsRepo,
		grRepo: grRepo,
		pRepo:  pRepo,
	}
}

type UseCase struct {
	gsRepo GameStateRepository
	grRepo GameRequestRepository
	pRepo  PlayerRepository
}

type GameRequestRepository interface {
	FindPending(challenger string) (*gamerequest.GameRequest, error)
	Save(gs *gamerequest.GameRequest) error
}

type PlayerRepository interface {
	Get(id string) (player.Player, error)
}

func (uc UseCase) Start(gob GameOutputBoundary, pId string) {
	// TODO matchmaking
	gr, err := uc.grRepo.FindPending(pId)
	if err != nil {
		// handle error
	}

	if gr != nil {
		gr.ChallengerId = pId
		gr.Pending = false

		// TODO transaction?
		uc.grRepo.Save(gr)

		// handle error
		pOne, _ := uc.pRepo.Get(gr.OwnerId)
		pTwo, _ := uc.pRepo.Get(gr.ChallengerId)

		gs := gamestate.New(uuid.NewString(), pOne, pTwo, []gamestate.History{}, pOne.Id, false)

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
