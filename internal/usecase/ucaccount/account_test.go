package ucaccount

import (
	"elton-okawa/battleship/internal/entity/account"
	"elton-okawa/battleship/internal/usecase/ucerror"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type MockDb struct {
	getError  bool
	saveError bool
	acc       account.Account
}

func (p *MockDb) Get(login string) (account.Account, error) {
	if p.getError {
		return account.Account{}, errors.New("Get mock error")
	}

	return p.acc, nil
}

func (p *MockDb) Save(acc account.Account) error {
	if p.saveError {
		return errors.New("Save mock error")
	}

	p.acc = acc

	return nil
}

type MockOutput struct {
	acc   AccountDTO
	login LoginDTO
}

func (out *MockOutput) CreateAccountResponse(acc AccountDTO) {
	out.acc = acc

}

func (out *MockOutput) LoginResponse(login LoginDTO) {
	out.login = login
}

func TestCreateAccount(t *testing.T) {
	assert := assert.New(t)

	username := "username"
	db := &MockDb{}
	out := &MockOutput{}

	useCase := New(db)
	err := useCase.CreateAccount(out, username, "password")

	assert.Nilf(err, "unexpected error %v", err)
	assert.Equal(username, db.acc.Login, "persisted account equal")
	assert.NotEmpty(db.acc.PasswordHash, "persisted account must have a password hash")
	assert.NotEmpty(db.acc.PlayerId, "persisted account must have a player id")

	assert.NotEmpty(out.acc.Id, "output id must not be empty")
	assert.Equal(username, out.acc.Login, "output login must be equal")
}

func TestCreateAccount_SaveError(t *testing.T) {
	assert := assert.New(t)

	db := &MockDb{saveError: true}
	out := &MockOutput{}

	useCase := New(db)
	err := useCase.CreateAccount(out, "username", "password")

	var e *ucerror.Error
	if assert.ErrorAs(err, &e, "use case error") {
		assert.Equal(e.Code, ucerror.GenericError, "use case error code")
	}
}

func TestLogin(t *testing.T) {
	assert := assert.New(t)

	username := "username"
	password := "password"
	acc, _ := account.New(username, password)
	db := &MockDb{acc: acc}
	out := &MockOutput{}

	useCase := New(db)
	err := useCase.Login(out, username, password)

	assert.Nilf(err, "unexpected error %v", err)
	assert.Equal(username, out.login.Login)

	assert.NotEmpty(out.login.Id, "id must not be empty")
	assert.NotEmpty(out.login.Token, "token must not be empty")
	assert.LessOrEqual(time.Now().Add(30*time.Minute).Unix(), out.login.ExpiresAt, "token must expire at least in 30 minutes")
	assert.GreaterOrEqual(time.Now().Add(120*time.Minute).Unix(), out.login.ExpiresAt, "token must not be valid for more than 2 hours")
}

func TestLogin_IncorrectUsername(t *testing.T) {
	assert := assert.New(t)
	db := &MockDb{getError: true}
	out := &MockOutput{}

	useCase := New(db)
	err := useCase.Login(out, "username", "password")

	var e *ucerror.Error
	if assert.ErrorAs(err, &e, "use case error") {
		assert.Equal(ucerror.IncorrectUsername, e.Code)
	}
}

func TestLogin_IncorrectPassword(t *testing.T) {
	assert := assert.New(t)

	username := "username"
	password := "password"
	acc, _ := account.New(username, password)
	db := &MockDb{acc: acc}
	out := &MockOutput{}

	useCase := New(db)
	err := useCase.Login(out, username, "another-password")

	var e *ucerror.Error
	if assert.ErrorAs(err, &e, "use case error") {
		assert.Equal(ucerror.IncorrectPassword, e.Code)
	}
}

func TestCreateAccountAndLogin(t *testing.T) {
	assert := assert.New(t)

	username := "username"
	password := "password"

	db := &MockDb{}
	useCase := New(db)
	outCreate := &MockOutput{}
	createError := useCase.CreateAccount(outCreate, username, password)

	assert.Nilf(createError, "unexpected error %v", createError)

	outLogin := &MockOutput{}
	loginError := useCase.Login(outLogin, username, password)

	assert.Nilf(loginError, "unexpected error %v", loginError)
}
