package routes

import (
	"ca-server/src/controllers"

	"github.com/gorilla/mux"
)

func SetupProfileRoutes(api *mux.Router) {
	api.HandleFunc("/profile", controllers.GetProfileInfo).Methods("POST")
	api.HandleFunc("/update-profile", controllers.UpdateProfileInfo).Methods("POST")
	api.HandleFunc("/change-password", controllers.UpdatePassword).Methods("POST")
}
