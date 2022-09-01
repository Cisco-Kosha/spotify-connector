package app

import (
	httpSwagger "github.com/swaggo/http-swagger"
)

func (a *App) initializeRoutes() {
	var apiV1 = "/api/v1"

	// specification routes
	a.Router.HandleFunc(apiV1+"/specification/list", a.listConnectorSpecification).Methods("GET", "OPTIONS")
	a.Router.HandleFunc(apiV1+"/specification/test", a.testConnectorSpecification).Methods("POST", "OPTIONS")

	// catalog search
	a.Router.HandleFunc(apiV1+"/search/{catalog}", a.searchCatalog).Methods("GET", "OPTIONS")

	// profile
	a.Router.HandleFunc(apiV1+"/profile/{profile}", a.profile).Methods("GET", "OPTIONS")

	// Swagger
	a.Router.PathPrefix("/docs").Handler(httpSwagger.WrapHandler)
}
