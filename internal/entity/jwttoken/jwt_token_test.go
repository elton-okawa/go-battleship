package jwttoken

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	assert := assert.New(t)

	token, expiresAt, err := New("username", "player-id")

	assert.Nilf(err, "unexpected error %v", err)
	assert.GreaterOrEqual(expiresAt, time.Now().Add(30*time.Minute).Unix(), "token should be valid for at least 30 min")
	assert.LessOrEqual(expiresAt, time.Now().Add(120*time.Minute).Unix(), "token should not be valid for more than 2 hours")
	assert.NotEmpty(token, "token should not be empty")
}

func TestNewAndValidate(t *testing.T) {
	assert := assert.New(t)

	token, _, err := New("username", "player-id")
	assert.Nilf(err, "unexpected error %v", err)

	claim, err := Validate(token)

	expected := Claim{
		Iss:    "Golang Battleship",
		Sub:    "username",
		Player: "player-id",
		Exp:    claim.Exp,
	}

	assert.Nilf(err, "unexpected error %v", err)
	assert.Equal(expected, claim)
}

func TestValidate_Expired(t *testing.T) {
	assert := assert.New(t)

	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJc3MiOiJHb2xhbmcgQmF0dGxlc2hpcCIsIlN1YiI6InVzZXJuYW1lIiwiRXhwIjoxNjUwMTMxMzM3fQ.pBH4X3A0N_LpAViwu56pmDOmKI2CXYCxXzaU08nGjNE"
	_, err := Validate(token)

	// jwt library does not return the error from Valid() function but
	// a new error using the same message
	assert.ErrorContains(err, "expired token")
}

// TODO think if this test is meaningful
func TestValidate_InvalidToken(t *testing.T) {
	assert := assert.New(t)

	token := "invalid"
	_, err := Validate(token)

	// jwt library return an error without message
	assert.NotNil(err)
}
