package routes

import (
	"github.com/gorilla/mux"
	"github.com/subhroacharjee/auth/lib/controller"
)

func HealthCheckHandler(r *mux.Router, c controller.ControllerResolver) {
	r.HandleFunc("/health-check", c.HealthCheck.HeathCheckHandler).Methods("GET")
}
