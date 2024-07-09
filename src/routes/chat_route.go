package routes

import (
	"ca-server/src/controllers"

	"github.com/gorilla/mux"
)

func SetupChatRoutes(r *mux.Router) {
	r.HandleFunc("/ws", controllers.HandleConnections).Methods("GET")
}
