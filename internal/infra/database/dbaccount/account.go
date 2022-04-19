package dbaccount

import (
	"elton-okawa/battleship/internal/entity/account"
	"elton-okawa/battleship/internal/infra/database"
)

type Account struct {
	Id           string `json:"id"`
	Login        string `json:"login"`
	PasswordHash string `json:"password"`
}

func (a *Account) GetId() string {
	return a.Id
}

func (a *Account) SetId(id string) {
	a.Id = id
}

type Dao struct {
	driver database.JsonDatabase
}

func New(filepath string) Dao {
	return Dao{
		driver: database.NewJsonDatabase(filepath),
	}
}

func (db Dao) Save(e account.Account) error {
	account := Account{
		Id:           e.Id,
		Login:        e.Login,
		PasswordHash: e.PasswordHash,
	}

	return db.driver.Save(&account)
}

func (db Dao) Get(login string) (account.Account, error) {
	var model Account
	err := db.driver.FindFirstBy("login", login, &model)

	acc := account.Account{
		Id:           model.Id,
		Login:        model.Login,
		PasswordHash: model.PasswordHash,
	}

	return acc, err
}
