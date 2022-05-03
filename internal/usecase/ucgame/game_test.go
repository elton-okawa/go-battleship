package ucgame

import (
	"elton-okawa/battleship/internal/entity"
	"elton-okawa/battleship/internal/entity/gamerequest"
	"elton-okawa/battleship/internal/entity/gamestate"
	"elton-okawa/battleship/internal/usecase/ucerror"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockGameStateRepo struct {
	saveData  gamestate.GameState
	saveError bool
}

func (r *MockGameStateRepo) Save(gs gamestate.GameState) error {
	if r.saveError {
		return errors.New("save game state error")
	}

	r.saveData = gs
	return nil
}

func (r *MockGameStateRepo) Get(id string) (*gamestate.GameState, error) {
	return nil, nil
}

type MockGameRequestRepo struct {
	ownNotFound      bool
	ownReadError     bool
	pendingNotFound  bool
	pendingReadError bool
	saveError        bool
	saveData         *gamerequest.GameRequest
	getData          gamerequest.GameRequest
}

func (r *MockGameRequestRepo) FindOwn(owner string) (*gamerequest.GameRequest, error) {
	if r.ownNotFound {
		return nil, entity.ErrNotFound
	}

	if r.ownReadError {
		return nil, errors.New("read own game request error")
	}

	return &r.getData, nil
}

func (r *MockGameRequestRepo) FindPending() (*gamerequest.GameRequest, error) {
	if r.pendingNotFound {
		return nil, entity.ErrNotFound
	}

	if r.pendingReadError {
		return nil, errors.New("read pending game request error")
	}

	// data := r.getData
	return &r.getData, nil
}

func (r *MockGameRequestRepo) Save(gs *gamerequest.GameRequest) error {
	if r.saveError {
		return errors.New("save game request error")
	}

	r.saveData = gs
	return nil
}

type MockOutput struct{}

func (mo *MockOutput) StartResult() {
}

func (mo *MockOutput) ShootResult(*gamestate.GameState, bool, int, error) {
}

func TestStart_Owner(t *testing.T) {
	assert := assert.New(t)

	gsRepo := &MockGameStateRepo{}
	grRepo := &MockGameRequestRepo{
		ownNotFound:     true,
		pendingNotFound: true,
	}
	out := &MockOutput{}
	pId := "player-id"

	useCase := New(gsRepo, grRepo)

	err := useCase.Start(out, pId)
	assert.Nilf(err, "unexpected error %v", err)
	assert.NotEmpty(grRepo.saveData.Id)
	assert.Equal(pId, grRepo.saveData.OwnerId)
	assert.Equal("", grRepo.saveData.ChallengerId)
	assert.True(grRepo.saveData.Pending, "game request must be pending")
}

func TestStart_Challenger(t *testing.T) {
	assert := assert.New(t)

	grFixture := gamerequest.New(
		"game-request-id",
		"other-id",
		"",
		true,
	)
	gsRepo := &MockGameStateRepo{}
	grRepo := &MockGameRequestRepo{
		ownNotFound:     true,
		pendingNotFound: false,
		getData:         grFixture,
	}
	out := &MockOutput{}
	pId := "player-id"

	useCase := New(gsRepo, grRepo)

	err := useCase.Start(out, pId)
	assert.Nilf(err, "unexpected error %v", err)
	assert.NotEmpty(grRepo.saveData.Id)
	assert.Equal(grFixture.OwnerId, grRepo.saveData.OwnerId)
	assert.Equal(pId, grRepo.saveData.ChallengerId)
	assert.False(grRepo.saveData.Pending, "game request must not be pending")

	assert.NotEmpty(gsRepo.saveData.Id)
	assert.Equal(grFixture.OwnerId, gsRepo.saveData.AccountOneId)
	assert.Equal(pId, gsRepo.saveData.AccountTwoId)
	assert.Empty(gsRepo.saveData.History, "new game state should have empty history")
	assert.Equal(grFixture.OwnerId, gsRepo.saveData.PlayerTurnId)
	assert.False(gsRepo.saveData.Finished, "new game state must not start finished")
	// TODO test configurable board size and number of ship
}

func TestStart_OwnGameRequestExist(t *testing.T) {
	assert := assert.New(t)

	gsRepo := &MockGameStateRepo{}
	grRepo := &MockGameRequestRepo{
		ownNotFound:     false,
		pendingNotFound: true,
	}
	out := &MockOutput{}
	pId := "player-id"

	useCase := New(gsRepo, grRepo)

	err := useCase.Start(out, pId)
	var e *ucerror.Error
	if assert.ErrorAs(err, &e, "use case error") {
		assert.Equal(ucerror.ExistingGameRequest, e.Code)
	}
}

func TestStart_OwnGameRequestReadError(t *testing.T) {
	assert := assert.New(t)

	gsRepo := &MockGameStateRepo{}
	grRepo := &MockGameRequestRepo{
		ownReadError:    true,
		pendingNotFound: true,
	}
	out := &MockOutput{}
	pId := "player-id"

	useCase := New(gsRepo, grRepo)

	err := useCase.Start(out, pId)
	var e *ucerror.Error
	if assert.ErrorAs(err, &e, "use case error") {
		assert.Equal(ucerror.ServerError, e.Code)
	}
}

func TestStart_PendingGameRequestReadError(t *testing.T) {
	assert := assert.New(t)

	gsRepo := &MockGameStateRepo{}
	grRepo := &MockGameRequestRepo{
		ownNotFound:      true,
		pendingReadError: true,
	}
	out := &MockOutput{}
	pId := "player-id"

	useCase := New(gsRepo, grRepo)

	err := useCase.Start(out, pId)

	var e *ucerror.Error
	if assert.ErrorAs(err, &e) {
		assert.Equal(ucerror.ServerError, e.Code)
	}
}

func TestStart_OwnGameRequestSaveError(t *testing.T) {
	assert := assert.New(t)

	gsRepo := &MockGameStateRepo{}
	grRepo := &MockGameRequestRepo{
		ownNotFound:     true,
		pendingNotFound: true,
		saveError:       true,
	}
	out := &MockOutput{}
	pId := "player-id"

	useCase := New(gsRepo, grRepo)

	err := useCase.Start(out, pId)

	var e *ucerror.Error
	if assert.ErrorAs(err, &e) {
		assert.Equal(ucerror.ServerError, e.Code)
	}
}

func TestStart_ChallengerGameRequestSaveError(t *testing.T) {
	assert := assert.New(t)

	gsRepo := &MockGameStateRepo{}
	grRepo := &MockGameRequestRepo{
		ownNotFound:     true,
		pendingNotFound: false,
		saveError:       true,
	}
	out := &MockOutput{}
	pId := "player-id"

	useCase := New(gsRepo, grRepo)

	err := useCase.Start(out, pId)

	var e *ucerror.Error
	if assert.ErrorAs(err, &e) {
		assert.Equal(ucerror.ServerError, e.Code)
	}
}

func TestStart_ChallengerGameStateSaveError(t *testing.T) {
	assert := assert.New(t)

	gsRepo := &MockGameStateRepo{
		saveError: true,
	}
	grRepo := &MockGameRequestRepo{
		ownNotFound:     true,
		pendingNotFound: false,
	}
	out := &MockOutput{}
	pId := "player-id"

	useCase := New(gsRepo, grRepo)

	err := useCase.Start(out, pId)

	var e *ucerror.Error
	if assert.ErrorAs(err, &e) {
		assert.Equal(ucerror.ServerError, e.Code)
	}
}
