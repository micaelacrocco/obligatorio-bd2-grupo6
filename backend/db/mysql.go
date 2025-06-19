package db

import (
	"EleccionesUcu/configuration"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func ConnectDb() *sql.DB {
	configuration.Init()

	database, err := sql.Open("mysql", configuration.GetDSN())
	if err != nil {
		panic(err.Error())
	}
	err = database.Ping()
	if err != nil {
		panic(err.Error())
	}
	return database
}
