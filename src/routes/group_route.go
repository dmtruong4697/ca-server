package routes

import (
	"ca-server/src/controllers"

	"github.com/gorilla/mux"
)

func SetupGroupRoutes(api *mux.Router) {
	api.HandleFunc("/group-info", controllers.GetGroupInfo).Methods("POST")
	api.HandleFunc("/group-list", controllers.GetGroupList).Methods("POST")
}
