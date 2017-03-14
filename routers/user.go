package routers

import (
	"github.com/gorilla/mux"
	"taskmanager/controllers"
)

func SetUserRoutes(router *mux.Router) *mux.Router  {
	router.HandleFunc("/users/register",controllers.Regeister).Methods("POST")
	router.HandleFunc("/users/login",controllers.Login).Methods("POST")
	return router
}
