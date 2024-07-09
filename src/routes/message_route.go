package routes

import (
	"ca-server/src/controllers"

	"github.com/gorilla/mux"
)

func SetupMessageRoutes(api *mux.Router) {
	api.HandleFunc("/messages", controllers.GetGroupMessage).Methods("POST")
}
