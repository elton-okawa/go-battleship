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
	acc       account.Account
	err       error
	token     string
	expiresAt int64
}

func (out *MockOutput) CreateAccountResponse(acc account.Account, err error) {
	out.acc = acc
	out.err = err
}

func (out *MockOutput) LoginResponse(acc account.Account, token string, expiresAt int64, err error) {
	out.acc = acc
	out.err = err
	out.expiresAt = expiresAt
	out.token = token
}

func TestCreateAccount(t *testing.T) {
	assert := assert.New(t)

	username := "username"
	db := &MockDb{}
	out := &MockOutput{}

	useCase := New(db)
	useCase.CreateAccount(out, username, "password")

	assert.Nilf(out.err, "unexpected error %v", out.err)
	assert.Equal(username, db.acc.Login, "persisted account equal")
	assert.NotEmpty(db.acc.PasswordHash, "persisted account must have a password hash")
	assert.Equal(username, out.acc.Login, "output login must be equal")
	assert.NotEmpty(out.acc.PasswordHash, "output password hash must not be empty")
}

func TestCreateAccountSaveError(t *testing.T) {
	assert := assert.New(t)

	db := &MockDb{saveError: true}
	out := &MockOutput{}

	useCase := New(db)
	useCase.CreateAccount(out, "username", "password")

	var e *ucerror.Error
	if assert.ErrorAs(out.err, &e, "use case error") {
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
	useCase.Login(out, username, password)

	assert.Nilf(out.err, "unexpected error %v", out.err)
	assert.Equal(username, out.acc.Login)
	assert.Equal(acc.PasswordHash, out.acc.PasswordHash)
	assert.NotEmpty(out.token, "token must not be empty")
	assert.LessOrEqual(time.Now().Add(30*time.Minute).Unix(), out.expiresAt, "token must expire at least in 30 minutes")
	assert.GreaterOrEqual(time.Now().Add(120*time.Minute).Unix(), out.expiresAt, "token must not be valid for more than 2 hours")
}

func TestLoginIncorrectUsername(t *testing.T) {
	assert := assert.New(t)
	db := &MockDb{getError: true}
	out := &MockOutput{}

	useCase := New(db)
	useCase.Login(out, "username", "password")

	var e *ucerror.Error
	if assert.ErrorAs(out.err, &e, "use case error") {
		assert.Equal(ucerror.IncorrectUsername, e.Code)
	}
}

func TestLoginIncorrectPassword(t *testing.T) {
	assert := assert.New(t)

	username := "username"
	password := "password"
	acc, _ := account.New(username, password)
	db := &MockDb{acc: acc}
	out := &MockOutput{}

	useCase := New(db)
	useCase.Login(out, username, "another-password")

	var e *ucerror.Error
	if assert.ErrorAs(out.err, &e, "use case error") {
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
	useCase.CreateAccount(outCreate, username, password)

	assert.Nilf(outCreate.err, "unexpected error %v", outCreate.err)

	outLogin := &MockOutput{}
	useCase.Login(outLogin, username, password)

	assert.Nilf(outLogin.err, "unexpected error %v", outLogin.err)
}
