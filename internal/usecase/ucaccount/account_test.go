package ucaccount

import (
	"elton-okawa/battleship/internal/entity/account"
	"elton-okawa/battleship/internal/usecase/ucerror"
	"errors"
	"testing"
	"time"
)

type MockPersistence struct {
	getError  bool
	saveError bool
	acc       account.Account
}

func (p *MockPersistence) Get(login string) (account.Account, error) {
	if p.getError {
		return account.Account{}, errors.New("Get mock error")
	}

	return p.acc, nil
}

func (p *MockPersistence) Save(acc account.Account) error {
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
	username := "username"
	persistence := &MockPersistence{}
	out := &MockOutput{}

	useCase := New(persistence)
	useCase.CreateAccount(out, username, "password")

	if out.err != nil {
		t.Fatalf("Unexpected error %v", out.err)
	}

	if persistence.acc.Login != username {
		t.Errorf("Persisted account saved with different login (expected: %s, actual: %s)", persistence.acc.Login, username)
	}
	if persistence.acc.PasswordHash == "" {
		t.Errorf("Persisted account must have a password hash")
	}

	if out.acc.Login != username {
		t.Errorf("Output account with different login than expected (expected: %s, actual: %s)", out.acc.Login, username)
	}
	if out.acc.PasswordHash == "" {
		t.Errorf("Output account must have a password hash")
	}
}

func TestCreateAccountSaveError(t *testing.T) {
	persistence := &MockPersistence{saveError: true}
	output := &MockOutput{}

	useCase := New(persistence)
	useCase.CreateAccount(output, "username", "password")

	var e *ucerror.Error
	if errors.As(output.err, &e) {
		if e.Code != ucerror.GenericError {
			t.Errorf("Expected %d, got %d", ucerror.GenericError, e.Code)
		}
	} else {
		t.Errorf("Expected ucerror.Error type %v", output.err)
	}
}

func TestLogin(t *testing.T) {
	username := "username"
	password := "password"
	acc, _ := account.New(username, password)
	db := &MockPersistence{acc: acc}
	out := &MockOutput{}

	useCase := New(db)
	useCase.Login(out, username, password)

	if out.err != nil {
		t.Fatalf("Unexpected error %v", out.err)
	}

	if out.acc.Login != username {
		t.Errorf("Login should be equal username (expected: %s, actual: %s)", out.acc.Login, username)
	}
	if out.acc.PasswordHash != acc.PasswordHash {
		t.Errorf("PasswordHash should be equal (expected: %s, actual: %s)", out.acc.PasswordHash, acc.PasswordHash)
	}
	if out.token == "" {
		t.Error("Token must not be empty", out.token)
	}
	if out.expiresAt <= time.Now().Add(30*time.Minute).Unix() {
		t.Errorf("Token should be valid at least for 30 min")
	}
}

func TestLoginIncorrectUsername(t *testing.T) {
	db := &MockPersistence{getError: true}
	out := &MockOutput{}

	useCase := New(db)
	useCase.Login(out, "username", "password")

	var e *ucerror.Error
	if errors.As(out.err, &e) {
		if e.Code != ucerror.IncorrectUsername {
			t.Errorf("Expected: %d, Got: %d", ucerror.IncorrectUsername, e.Code)
		}
	} else {
		t.Fatalf("Unexpected error %v", out.err)
	}
}

func TestLoginIncorrectPassword(t *testing.T) {
	username := "username"
	password := "password"
	acc, _ := account.New(username, password)
	db := &MockPersistence{acc: acc}
	out := &MockOutput{}

	useCase := New(db)
	useCase.Login(out, username, "another-password")

	var e *ucerror.Error
	if errors.As(out.err, &e) {
		if e.Code != ucerror.IncorrectPassword {
			t.Errorf("Expected code: %d, got: %d", ucerror.IncorrectPassword, e.Code)
		}
	} else {
		t.Fatalf("Unexpected error %v", out.err)
	}
}

func TestCreateAccountAndLogin(t *testing.T) {
	username := "username"
	password := "password"

	db := &MockPersistence{}
	useCase := New(db)
	outCreate := &MockOutput{}
	useCase.CreateAccount(outCreate, username, password)

	if outCreate.err != nil {
		t.Fatalf("Unexpected error %v", outCreate.err)
	}

	outLogin := &MockOutput{}
	useCase.Login(outLogin, username, password)

	if outLogin.err != nil {
		t.Fatalf("Unexpected error %v", outLogin.err)
	}
}
