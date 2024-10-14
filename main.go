package main

import (
	"net/http"

	"github.com/tirlochanarora16/blogging_api_go/routes"
)

func main() {
	http.HandleFunc("/posts", routes.HandleRoutes)

	http.ListenAndServe(":3000", nil)
}
