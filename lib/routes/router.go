package routes

import (
	"github.com/gorilla/mux"
	"github.com/subhroacharjee/auth/lib/controller"
	"github.com/subhroacharjee/auth/lib/middleware"
)

func NewRouter(c controller.ControllerResolver) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	subrouter := router.PathPrefix("/").Subrouter()
	subrouter.Use(middleware.JSONMiddleware)
	HealthCheckHandler(subrouter, c)
	UserRoutesHandler(subrouter, c)
	return router
}
