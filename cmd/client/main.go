package main

import (
	"flag"
	"net/http"
	_ "userService/pkg/dataBase/mySql/migrations"
)

func main() {
	address := flag.String(
		"address",
		":8001",
		"Port of this web service")
	//dbConnection := flag.String(
	//	"dbConnection",
	//	"root:root@/users?parseTime=true",
	//	"mySQL connection string")
	flag.Parse()

	app := Application{
		ErrorLog: ErrorLog,
		InfoLog:  InfoLog,
	}

	srv := &http.Server{
		Addr:     *address,
		ErrorLog: ErrorLog,
		Handler:  app.routes(),
	}

	InfoLog.Printf("Server started on %s", *address)
	err := srv.ListenAndServe()
	ErrorLog.Fatal(err)
}
