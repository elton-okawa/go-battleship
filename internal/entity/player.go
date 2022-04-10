package entity

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Player struct {
	Id           string
	passwordHash string
}

// TODO we might want to impose some password restrictions like length, characters
func NewPlayer(password string) (Player, error) {
	saltedBytes := []byte(password)
	hashedBytes, err := bcrypt.GenerateFromPassword(saltedBytes, bcrypt.DefaultCost)
	if err != nil {
		return Player{}, err
	}

	player := Player{
		Id:           uuid.NewString(),
		passwordHash: string(hashedBytes),
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
	return bcrypt.CompareHashAndPassword([]byte(p.passwordHash), []byte(password))
}
