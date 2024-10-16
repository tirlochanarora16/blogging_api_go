package main

import (
	"net/http"

	_ "github.com/lib/pq"

	"github.com/tirlochanarora16/blogging_api_go/db"
	"github.com/tirlochanarora16/blogging_api_go/routes"
)

func main() {

	connStr := "user=tirlochan password=password dbname=blog_api sslmode=disable"

	db.InitDB(connStr)

	http.HandleFunc("/posts", routes.HandleRoutes)

	http.ListenAndServe(":3000", nil)
}
