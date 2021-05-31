package main

import (
	"log"
	"os"
	"userService/pkg/dataBase/mySql"

	_ "github.com/go-sql-driver/mysql"
)

var InfoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.LUTC)
var ErrorLog = log.New(os.Stderr, "ERROR\t", log.Ldate|log.LUTC|log.Lshortfile)

type Application struct {
	ErrorLog *log.Logger
	InfoLog  *log.Logger
	Db       *mySql.UserModel
}
