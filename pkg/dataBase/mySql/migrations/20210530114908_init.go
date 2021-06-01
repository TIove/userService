package migrations

import (
	"database/sql"
	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upInit, downInit) // TODO add goose
}

func upInit(tx *sql.Tx) error {
	_, err := tx.Exec("CREATE TABLE users (id BINARY(36) NOT NULL PRIMARY KEY, name TEXT NOT NULL, birthday DATETIME NOT NULL)")
	if err != nil {
		return err
	}
	return nil
}

func downInit(tx *sql.Tx) error {
	_, err := tx.Exec("DROP TABLE users")
	if err != nil {
		return err
	}
	return nil
}
