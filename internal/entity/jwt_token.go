package entity

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

type TokenClaim struct {
	Username  string
	ExpiresAt int64
}

// TODO retrieve it from a secret storage
var hmacKey = []byte("super-secret")
var ErrExpiredToken = errors.New("expired token")

func NewJwtToken(username string) (string, int64, error) {
	expires := time.Now().Add(1 * time.Hour).Unix()

	claims := TokenClaim{
		Username:  username,
		ExpiresAt: expires,
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

func (tc TokenClaim) Valid() error {
	if tc.ExpiresAt >= time.Now().Unix() {
		return ErrExpiredToken
	}

	return nil
}

func ValidateJwtToken(tokenStr string) (TokenClaim, error) {
	claims := TokenClaim{}
	token, err := jwt.ParseWithClaims(tokenStr, &claims, func(token *jwt.Token) (interface{}, error) {
		return hmacKey, nil
	})

	if err == nil && token.Valid {
		return claims, nil
	} else {
		return TokenClaim{}, err
	}
}
