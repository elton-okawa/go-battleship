package entity

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Player struct {
	Id           string
	Login        string
	PasswordHash string
}

func NewPlayer(login, password string) (Player, error) {
	saltedBytes := []byte(password)
	hashedBytes, err := bcrypt.GenerateFromPassword(saltedBytes, bcrypt.DefaultCost)
	if err != nil {
		return Player{}, err
	}

	player := Player{
		Id:           uuid.NewString(),
		Login:        login,
		PasswordHash: string(hashedBytes),
	}

	return player, nil
}

func (p *Player) GetId() string {
	return p.Id
}

func (p *Player) SetId(id string) {
	p.Id = id
}

func (p *Player) Authenticate(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(p.PasswordHash), []byte(password))
}
