package main

import (
	"taskmanager/common"
	"taskmanager/routers"
	"github.com/codegangsta/negroni"
	"net/http"
	"log"
)

func main()  {
	common.StartUp()
	router:=routers.InitRoutes()
	n:=negroni.Classic()
	n.UseHandler(router)

	server:=&http.Server{
		Addr:common.AppConfig.Server,
		Handler:n,
	}
	log.Println("Listening")
	server.ListenAndServe()
}
