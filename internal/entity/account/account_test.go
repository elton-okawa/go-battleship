package account

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAccount_UniqueId(t *testing.T) {
	assert := assert.New(t)

	playerOne, errOne := New("player-one", "secret")
	playerTwo, errTwo := New("player-two", "password")

	assert.Nilf(errOne, "unexpected error %v", errOne)
	assert.Nilf(errTwo, "unexpected error %v", errTwo)
	assert.NotEqual(playerOne.GetId(), playerTwo.GetId(), "id must be unique")
}

func TestNewAccount_HashedPassword(t *testing.T) {
	assert := assert.New(t)
	password := "secret"

	player, _ := New("player", password)

	assert.NotEqual(password, player.PasswordHash, "not store plain password")
}

func TestNewAccount_SaltedPassword(t *testing.T) {
	assert := assert.New(t)
	password := "secret"

	playerOne, _ := New("player-one", password)
	playerTwo, _ := New("player-two", password)

	assert.NotEqual(playerOne.PasswordHash, playerTwo.PasswordHash, "same password must not be stored with the same value")
}

func TestAuthentication(t *testing.T) {
	assert := assert.New(t)
	password := "secret"

	account, _ := New("player", password)

	assert.Nil(account.Authenticate(password), "should authenticate without problems")
}

func TestAuthentication_WrongPassword(t *testing.T) {
	assert := assert.New(t)
	password := "secret"

	account, _ := New("player", password)

	assert.NotNil(account.Authenticate("another"), "must not be able to authenticate with wrong password")
}
