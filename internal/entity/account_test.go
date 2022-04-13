package entity

import "testing"

func TestNewAccountUniqueId(t *testing.T) {
	playerOne, errOne := NewAccount("player-one", "secret")
	playerTwo, errTwo := NewAccount("player-two", "password")

	if errOne != nil || errTwo != nil {
		t.Errorf("Error while creating new player one")
	}

	if playerOne.GetId() == playerTwo.GetId() {
		t.Errorf("NewAccount id must be unique")
	}
}

func TestNewAccountHashedPassword(t *testing.T) {
	password := "secret"

	player, _ := NewAccount("player", password)

	if player.PasswordHash == password {
		t.Error("We must not store plain password")
	}
}

func TestNewAccountSaltedPassword(t *testing.T) {
	password := "secret"

	playerOne, _ := NewAccount("player-one", password)
	playerTwo, _ := NewAccount("player-two", password)

	if playerOne.PasswordHash == playerTwo.PasswordHash {
		t.Error("Players with same password have the same salted hashed password")
	}
}

func TestAccountAuthentication(t *testing.T) {
	password := "secret"

	player, _ := NewAccount("player", password)

	if player.Authenticate("another") == nil {
		t.Error("Player must not be able to authenticate using other password")
	}

	if player.Authenticate(password) != nil {
		t.Error("Player must be able to authenticate using created password")
	}
}
