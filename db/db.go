package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
)

var DB *sql.DB

func InitDB(connStr string, runMigration bool) {
	var err error
	DB, err = sql.Open("postgres", connStr)

	if err != nil {
		panic(err)
	}

	err = DB.Ping()

	if err != nil {
		panic(err)
	}

	if runMigration {
		runDBMigration()
	}

	fmt.Println("Successfully connected to DB!")
}

func runDBMigration() {
	files, err := os.ReadDir("migrations")

	if err != nil {
		log.Fatal("Error reading the files in the migration directory")
	}

	lastSqlfile := files[len(files)-1]

	fileLocation := []string{"migrations/", lastSqlfile.Name()}

	fileContent, err := os.ReadFile(strings.Join(fileLocation, ""))

	if err != nil {
		log.Fatal("Error reading SQL file content from the migrations directory")
	}

	_, err = DB.Exec(string(fileContent))

	if err != nil {
		log.Fatal("Error in applying migration from latest file", err)
	}

	fmt.Println("Successfully applied migtation to the DB from the file: ", fileLocation)
}
