package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	DB_DRIVER   = "postgres"
	DB_HOST     = "postgres"
	DB_PORT     = 5432
	DB_USER     = "postgres"
	DB_PASSWORD = "postgres"
	DB_NAME     = "postgres"
)

func getDatabase() *sql.DB {
	dbinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME)
	postgres, err := sql.Open(DB_DRIVER, dbinfo)
	checkErr(err)
	return postgres
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
