package mySql

import (
	"database/sql"
	"github.com/google/uuid"
	"userService/pkg/models/dbModels"
)

type UserModel struct {
	Db *sql.DB
}

func (model *UserModel) Insert(dbUser dbModels.User) (sql.Result, error) {
	query := "INSERT INTO users VALUES (?, ?, ?)"

	res, err := model.Db.Exec(query, dbUser.Id, dbUser.Name, dbUser.Birthday)
	return res, err
}

func (model *UserModel) Get(id uuid.UUID) (*dbModels.User, error) {
	query := "select * from users where id = ? LIMIT 1"

	row := model.Db.QueryRow(query, id)

	result := dbModels.User{}
	err := row.Scan(&result.Id, &result.Name, &result.Birthday)

	return &result, err
}

func (model *UserModel) Delete(id uuid.UUID) (bool, error) {
	query := "delete from users where id = ?"

	_, err := model.Db.Exec(query, id)

	if err != nil {
		return false, err
	}

	return true, nil
}
