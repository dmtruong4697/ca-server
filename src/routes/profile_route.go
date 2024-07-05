package routes

import (
	"ca-server/src/controllers"

	"github.com/gorilla/mux"
)

func SetupProfileRoutes(r *mux.Router) {
	r.HandleFunc("/register", controllers.Register).Methods("POST")
	r.HandleFunc("/validate-email", controllers.ValidateEmail).Methods("POST")
	r.HandleFunc("/login", controllers.Login).Methods("POST")
	r.HandleFunc("/logout", controllers.Logout).Methods("POST")
}
