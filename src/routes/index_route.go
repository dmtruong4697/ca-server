package routes

import (
	"ca-server/src/middlewares"
	_ "go-authentication/docs"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

func SetupRouter() *mux.Router {

	r := mux.NewRouter()

	// Protected routes
	api := r.PathPrefix("/api").Subrouter()
	api.Use(middlewares.AuthMiddleware)

	// Swagger route
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	// Auth routes
	SetupAuthRoutes(r)

	// User routes
	SetupUserRoutes(api)

	// Chat routes
	SetupChatRoutes(r)

	// Channel routes
	SetupChannelRoutes(api)

	return r
}
