package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"github.com/tirlochanarora16/blogging_api_go/db"
	"github.com/tirlochanarora16/blogging_api_go/routes"
)

func main() {
	migrate := flag.Bool("migrate", false, "Use migrations")

	flag.Parse()

	if _, err := os.Stat(".env"); err == nil {
		err := godotenv.Load()
		if err != nil {
			log.Println("Error loading .env file, but continuing with existing environment variables")
		}
	}

	connStr := os.Getenv("DB_CONN_STR")
	if connStr == "" {
		log.Fatal("DB_CONN_STR is not set")
	}

	db.InitDB(connStr, *migrate)

	http.HandleFunc("/posts", routes.HandleRoutes)

	err := http.ListenAndServe(":3000", nil)

	if err != nil {
		log.Fatal("Error starting the server", err)
	}
}
