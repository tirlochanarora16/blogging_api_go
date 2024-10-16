package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/tirlochanarora16/blogging_api_go/routes"
)

func main() {

	connStr := "user=tirlochan password=password dbname=blog_api sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	http.HandleFunc("/posts", routes.HandleRoutes)

	http.ListenAndServe(":3000", nil)
}
