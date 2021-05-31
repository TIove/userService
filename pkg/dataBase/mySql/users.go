package mySql

import (
	"database/sql"
	"github.com/google/uuid"
	"time"
	"userService/pkg/models/dbModels"
)

type UserModel struct {
	Db *sql.DB
}

func (model *UserModel) Insert(dbUser dbModels.User) (uuid.UUID, error) {
	return uuid.New(), nil //TODO add logic to db methods
}

func (model *UserModel) Get(id int) (*dbModels.User, error) {
	return &dbModels.User{Id: uuid.New(), Name: "Name", Birthday: time.Now()}, nil
}

func (model *UserModel) Delete(id uuid.UUID) (bool, error) {
	return true, nil
}
