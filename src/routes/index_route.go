package routes

import (
	"ca-server/src/middlewares"

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

	// Profile routes
	SetupProfileRoutes(api)

	// User routes
	SetupUserRoutes(api)

	// Relationship routes
	SetupRelationshipRoutes(api)

	// Group routes
	SetupGroupRoutes(api)

	return r
}
