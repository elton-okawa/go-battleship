package ucgame

import (
	"elton-okawa/battleship/internal/entity"
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
	FindOwn(owner string) (gamerequest.GameRequest, error)
	FindPending() (gamerequest.GameRequest, error)
	Save(gs gamerequest.GameRequest) error
}

type GameStateRepository interface {
	Save(gamestate.GameState) error
	Get(string) (gamestate.GameState, error)
}

type GameOutputBoundary interface {
	StartResult()
	ShootResult(*gamestate.GameState, bool, int, error)
}

func (uc UseCase) Start(gob GameOutputBoundary, pId string) error {
	_, ownReqErr := uc.grRepo.FindOwn(pId)
	if ownReqErr == nil {
		useCaseError := ucerror.New(
			"cannot create a new game while you already have one waiting for an opponent",
			ucerror.ExistingGameRequest,
			nil,
		)
		return useCaseError
	} else if !errors.Is(ownReqErr, entity.ErrNotFound) {
		useCaseError := ucerror.New(
			"error while finding own game request",
			ucerror.ServerError,
			ownReqErr,
		)
		return useCaseError
	}

	// TODO better matchmaking
	gr, err := uc.grRepo.FindPending()
	if err == nil {
		if err := uc.fulfillGameRequest(gr, pId); err != nil {
			return err
		}
	} else if errors.Is(err, entity.ErrNotFound) {
		gr := gamerequest.New(uuid.NewString(), pId, "", true)
		if err := uc.grRepo.Save(gr); err != nil {
			saveGRequestErr := ucerror.New(
				"error while saving game request",
				ucerror.ServerError,
				err,
			)
			return saveGRequestErr
		}
	} else {
		useCaseError := ucerror.New(
			"error while reading game request",
			ucerror.ServerError,
			err,
		)
		return useCaseError
	}

	gob.StartResult()
	return nil
}

func (uc UseCase) fulfillGameRequest(gr gamerequest.GameRequest, pId string) error {
	gr.ChallengerId = pId
	gr.Pending = false

	// TODO transaction?
	if err := uc.grRepo.Save(gr); err != nil {
		saveGRequestErr := ucerror.New(
			"error while saving game request",
			ucerror.ServerError,
			err,
		)
		return saveGRequestErr
	}

	gs := gamestate.New(
		uuid.NewString(),
		gr.OwnerId,
		gr.ChallengerId,
		board.Initialize(8, 3),
		board.Initialize(8, 3),
		[]gamestate.Turn{},
		gr.OwnerId,
		false,
	)

	if err := uc.gsRepo.Save(gs); err != nil {
		saveGStateErr := ucerror.New(
			"error while saving game state",
			ucerror.ServerError,
			err,
		)
		return saveGStateErr
	}

	return nil
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
