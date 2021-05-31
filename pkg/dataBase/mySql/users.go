package mySql

import (
	"database/sql"
	"time"
	"userService/pkg/models/dbModels"
)

type UserModel struct {
	Db *sql.DB
}

func (model *UserModel) Insert(dbUser dbModels.User) (int, error) {
	return 1, nil //TODO add logic to db methods
}

func (model *UserModel) Get(id int) (*dbModels.User, error) {
	return &dbModels.User{Id: 1, Name: "Name", Birthday: time.Now()}, nil
}

func (model *UserModel) Delete(id int) (bool, error) {
	return true, nil
}
