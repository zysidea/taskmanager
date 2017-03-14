package routers

import (
	"github.com/gorilla/mux"
	"taskmanager/controllers"
	"github.com/codegangsta/negroni"
	"taskmanager/common"
)

func SetNoteRoutes(router *mux.Router) *mux.Router  {
	noteRouter:=mux.NewRouter()
	noteRouter.HandleFunc("/notes",controllers.CreateNote).Methods("POST")
	noteRouter.HandleFunc("/notes/{id}",controllers.UpdateNote).Methods("PUT")
	noteRouter.HandleFunc("/notes",controllers.GetNotes).Methods("GET")
	noteRouter.HandleFunc("notes/{id}",controllers.GetNoteById).Methods("GET")
	noteRouter.HandleFunc("notes/tasks/{id}",controllers.GetNoteByTask).Methods("GET")
	noteRouter.HandleFunc("notes/{id}",controllers.DeleteNote).Methods("DELETE")

	router.PathPrefix("/tasks").Handler(negroni.New(
		negroni.HandlerFunc(common.Authorize),
		negroni.Wrap(noteRouter),
	))

	return router
}
