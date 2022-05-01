package jwttoken

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

type Claim struct {
	Iss    string
	Sub    string
	Player string
	Exp    int64
}

// TODO retrieve it from a secret storage
var hmacKey = []byte("super-secret")
var ErrExpiredToken = errors.New("custom expired token message")

func New(username, pId string) (string, int64, error) {
	expires := time.Now().Add(1 * time.Hour).Unix()

	claims := Claim{
		Iss:    "Golang Battleship",
		Sub:    username,
		Player: pId,
		Exp:    expires,
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign
	if signed, err := token.SignedString(hmacKey); err == nil {
		return signed, expires, nil
	} else {
		return "", 0, err
	}
}

func (tc Claim) Valid() error {
	if tc.Exp <= time.Now().Unix() {
		return ErrExpiredToken
	}

	return nil
}

func Validate(tokenStr string) (Claim, error) {
	claims := Claim{}
	token, err := jwt.ParseWithClaims(tokenStr, &claims, func(token *jwt.Token) (interface{}, error) {
		return hmacKey, nil
	})

	if err == nil && token.Valid {
		return claims, nil
	} else {
		return Claim{}, err
	}
}
