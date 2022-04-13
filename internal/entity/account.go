package entity

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Account struct {
	Id           string
	Login        string
	PasswordHash string
}

func NewAccount(login, password string) (Account, error) {
	saltedBytes := []byte(password)
	hashedBytes, err := bcrypt.GenerateFromPassword(saltedBytes, bcrypt.DefaultCost)
	if err != nil {
		return Account{}, err
	}

	acc := Account{
		Id:           uuid.NewString(),
		Login:        login,
		PasswordHash: string(hashedBytes),
	}

	return acc, nil
}

func (p *Account) GetId() string {
	return p.Id
}

func (p *Account) SetId(id string) {
	p.Id = id
}

func (p *Account) Authenticate(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(p.PasswordHash), []byte(password))
}
