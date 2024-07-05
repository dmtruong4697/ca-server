package routes

import (
	"ca-server/src/controllers"

	"github.com/gorilla/mux"
)

func SetupAuthRoutes(r *mux.Router) {
	r.HandleFunc("/register", controllers.Register).Methods("POST")
	r.HandleFunc("/validate-email", controllers.ValidateEmail).Methods("POST")
}
