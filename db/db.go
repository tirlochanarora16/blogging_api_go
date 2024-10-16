package db

import (
	"database/sql"
	"fmt"
)

var DB *sql.DB

func InitDB(connStr string) {
	var err error
	DB, err = sql.Open("postgres", connStr)

	if err != nil {
		panic(err)
	}

	err = DB.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to DB!")
}
