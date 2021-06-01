package main

import (
	"database/sql"
	"flag"
	"net/http"
	"userService/pkg/dataBase/mySql"

	_ "github.com/go-sql-driver/mysql"
	_ "userService/pkg/dataBase/mySql/migrations"
)

func main() {
	address := flag.String(
		"address",
		":8001",
		"Port of this web service")
	dbConnection := flag.String(
		"dbConnection",
		"root:root@/?parseTime=true",
		"Database connection string")
	flag.Parse()

	db, err := openDb("mysql", *dbConnection)
	if err != nil {
		ErrorLog.Fatal(err)
	}
	defer db.Close()

	app := Application{
		ErrorLog: ErrorLog,
		InfoLog:  InfoLog,
		Db:       &mySql.UserModel{Db: db},
	}

	srv := &http.Server{
		Addr:     *address,
		ErrorLog: ErrorLog,
		Handler:  app.routes(),
	}

	InfoLog.Printf("Server started on %s", *address)
	err = srv.ListenAndServe()
	ErrorLog.Fatal(err)
}

func openDb(driverName string, dbCredentials string) (*sql.DB, error) {
	db, err := sql.Open(driverName, dbCredentials)

	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}

	createDb(db, "Users")

	return db, err
}

func createDb(db *sql.DB, dbName string) {
	_, err := db.Exec("CREATE DATABASE IF NOT EXISTS " + dbName)
	if err != nil {
		ErrorLog.Fatal(err)
	}

	_, err = db.Exec("USE " + dbName)
	if err != nil {
		ErrorLog.Fatal(err)
	}
}
