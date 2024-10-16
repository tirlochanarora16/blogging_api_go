package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"github.com/tirlochanarora16/blogging_api_go/db"
	"github.com/tirlochanarora16/blogging_api_go/routes"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading environment variables from .env file")
	}

	connStr := os.Getenv("DB_CONN_STR")

	db.InitDB(connStr)

	http.HandleFunc("/posts", routes.HandleRoutes)

	http.ListenAndServe(":3000", nil)
}
