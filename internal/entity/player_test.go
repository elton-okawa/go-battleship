package entity

import "testing"

func TestNewPlayerUniqueId(t *testing.T) {
	playerOne, errOne := NewPlayer("player-one", "secret")
	playerTwo, errTwo := NewPlayer("player-two", "password")

	if errOne != nil || errTwo != nil {
		t.Errorf("Error while creating new player one")
	}

	if playerOne.GetId() == playerTwo.GetId() {
		t.Errorf("NewPlayer id must be unique")
	}
}

func TestNewPlayerHashedPassword(t *testing.T) {
	password := "secret"

	player, _ := NewPlayer("player", password)

	if player.PasswordHash == password {
		t.Error("We must not store plain password")
	}
}

func TestNewPlayerSaltedPassword(t *testing.T) {
	password := "secret"

	playerOne, _ := NewPlayer("player-one", password)
	playerTwo, _ := NewPlayer("player-two", password)

	if playerOne.PasswordHash == playerTwo.PasswordHash {
		t.Error("Players with same password have the same salted hashed password")
	}
}

func TestPlayerAuthentication(t *testing.T) {
	password := "secret"

	player, _ := NewPlayer("player", password)

	if player.Authenticate("another") == nil {
		t.Error("Player must not be able to authenticate using other password")
	}

	if player.Authenticate(password) != nil {
		t.Error("Player must be able to authenticate using created password")
	}
}
