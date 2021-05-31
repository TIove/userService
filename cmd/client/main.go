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
		"root:root@/users?parseTime=true",
		"Database connection string")
	flag.Parse()

	db, err := openDb(*dbConnection, "mysql")
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

func openDb(dbConnection string, driverName string) (*sql.DB, error) {
	db, err := sql.Open(driverName, dbConnection)

	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, err
}
