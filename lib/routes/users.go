package routes

import (
	"github.com/gorilla/mux"
	"github.com/subhroacharjee/auth/lib/controller"
)

func UserRoutesHandler(r *mux.Router, c controller.ControllerResolver) {
	r.HandleFunc("/login", c.User.Login).Methods("POST")
	r.HandleFunc("/register", c.User.Register).Methods("POST")
}
