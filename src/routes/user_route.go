package routes

import (
	"ca-server/src/controllers"

	"github.com/gorilla/mux"
)

func SetupUserRoutes(api *mux.Router) {
	api.HandleFunc("/user-info", controllers.GetUserInfo).Methods("POST")
}
