package main

import (
	"ca-server/src/database"
	"ca-server/src/routes"
	"log"
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	// init database
	database.Connect()

	// Set up router
	r := routes.SetupRouter()

	// Swagger documentation route
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	// Start server
	log.Fatal(http.ListenAndServe(":8080", r))
}
