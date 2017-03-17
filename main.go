package main

import (
	"log"
	"net/http"
	"taskmanager/common"
	"taskmanager/routers"

	"github.com/codegangsta/negroni"
)

func main() {
	common.StartUp()
	router := routers.InitRoutes()
	n := negroni.Classic()
	n.UseHandler(router)

	server := &http.Server{
		Addr:    common.AppConfig.Server,
		Handler: n,
	}
	log.Println("Listening")
	server.ListenAndServe()
}
