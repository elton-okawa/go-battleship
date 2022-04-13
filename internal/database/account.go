package database

import (
	"elton-okawa/battleship/internal/entity"
)

type AccountModel struct {
	Id           string `json:"id"`
	Login        string `json:"login"`
	PasswordHash string `json:"password"`
}

func (a *AccountModel) GetId() string {
	return a.Id
}

func (a *AccountModel) SetId(id string) {
	a.Id = id
}

type AccountDao struct {
	driver JsonDatabase
}

func NewAccountDao(filepath string) AccountDao {
	return AccountDao{
		driver: NewJsonDatabase(filepath),
	}
}

func (db AccountDao) SaveAccount(e entity.Account) error {
	account := AccountModel{
		Id:           e.Id,
		Login:        e.Login,
		PasswordHash: e.PasswordHash,
	}

	return db.driver.Save(&account)
}

func (db AccountDao) GetAccount(login string) (entity.Account, error) {
	var model AccountModel
	err := db.driver.FindFirstBy("login", login, &model)

	acc := entity.Account{
		Id:           model.Id,
		Login:        model.Login,
		PasswordHash: model.PasswordHash,
	}

	return acc, err
}
