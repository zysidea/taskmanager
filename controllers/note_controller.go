package controllers

import (
	"encoding/json"
	"net/http"
	"taskmanager/common"
	"taskmanager/models"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
	"taskmanager/data"
)

//POST /notes
func CreateNote(w http.ResponseWriter, r *http.Request)  {
	var noteResource NoteResource
	err := json.NewDecoder(r.Body).Decode(&noteResource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid note data",
			http.StatusInternalServerError,
		)
	}
	if j, err := json.Marshal(noteResource); err != nil {
		common.DisplayUnexpectedAppError(w, err)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(j)
	}
}
//PUT /notes/{id}
func UpdateNote(w http.ResponseWriter, r *http.Request) {
	vars:=mux.Vars(r)
	id:=bson.ObjectIdHex(vars["id"])
	note := NoteResource{}
	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid note data",
			http.StatusInternalServerError,
		)
	}
	noteData:=&note.Data
	noteData.Id=id
	mc:=NewMongoContext()
	defer mc.Close()
	collection:=mc.GetCollection("notes")
	repo:=&data.NoteRepository{collection}
	if err:=repo.UpdateNote(noteData);err!=nil{
		common.DisplayUnexpectedAppError(w, err)
	} else {
		w.WriteHeader(http.StatusNoContent)
	}

}
//GET notes
func GetNotes(w http.ResponseWriter, r *http.Request) {

}
func GetNoteById(w http.ResponseWriter, r *http.Request) {

}
func GetNoteByTask(w http.ResponseWriter, r *http.Request) {

}
func DeleteNote(w http.ResponseWriter, r *http.Request) {

}
