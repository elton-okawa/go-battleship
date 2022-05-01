package account

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAccount_UniqueId(t *testing.T) {
	assert := assert.New(t)

	accOne, errOne := New("account-one", "secret")
	accTwo, errTwo := New("account-two", "password")

	assert.Nilf(errOne, "unexpected error %v", errOne)
	assert.Nilf(errTwo, "unexpected error %v", errTwo)
	assert.NotEqual(accOne.GetId(), accTwo.GetId(), "id must be unique")
}

func TestNewAccount_UniquePlayerId(t *testing.T) {
	assert := assert.New(t)

	accOne, errOne := New("account-one", "secret")
	accTwo, errTwo := New("account-two", "password")

	assert.Nilf(errOne, "unexpected error %v", errOne)
	assert.Nilf(errTwo, "unexpected error %v", errTwo)
	assert.NotEqual(accOne.PlayerId, accTwo.PlayerId, "player id must be unique")
}

func TestNewAccount_HashedPassword(t *testing.T) {
	assert := assert.New(t)
	password := "secret"

	acc, _ := New("account", password)

	assert.NotEqual(password, acc.PasswordHash, "not store plain password")
}

func TestNewAccount_SaltedPassword(t *testing.T) {
	assert := assert.New(t)
	password := "secret"

	accOne, _ := New("account-one", password)
	accTwo, _ := New("account-two", password)

	assert.NotEqual(accOne.PasswordHash, accTwo.PasswordHash, "same password must not be stored with the same value")
}

func TestAuthentication(t *testing.T) {
	assert := assert.New(t)
	password := "secret"

	acc, _ := New("account", password)

	assert.Nil(acc.Authenticate(password), "should authenticate without problems")
}

func TestAuthentication_WrongPassword(t *testing.T) {
	assert := assert.New(t)
	password := "secret"

	acc, _ := New("account", password)

	assert.NotNil(acc.Authenticate("another"), "must not be able to authenticate with wrong password")
}
