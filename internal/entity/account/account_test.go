package account

import "testing"

func TestNewAccountUniqueId(t *testing.T) {
	playerOne, errOne := New("player-one", "secret")
	playerTwo, errTwo := New("player-two", "password")

	if errOne != nil || errTwo != nil {
		t.Errorf("Error while creating new player one")
	}

	if playerOne.GetId() == playerTwo.GetId() {
		t.Errorf("NewAccount id must be unique")
	}
}

func TestNewAccountHashedPassword(t *testing.T) {
	password := "secret"

	player, _ := New("player", password)

	if player.PasswordHash == password {
		t.Error("We must not store plain password")
	}
}

func TestNewAccountSaltedPassword(t *testing.T) {
	password := "secret"

	playerOne, _ := New("player-one", password)
	playerTwo, _ := New("player-two", password)

	if playerOne.PasswordHash == playerTwo.PasswordHash {
		t.Error("Players with same password have the same salted hashed password")
	}
}

func TestAccountAuthentication(t *testing.T) {
	password := "secret"

	player, _ := New("player", password)

	if player.Authenticate("another") == nil {
		t.Error("Player must not be able to authenticate using other password")
	}

	if player.Authenticate(password) != nil {
		t.Error("Player must be able to authenticate using created password")
	}
}
