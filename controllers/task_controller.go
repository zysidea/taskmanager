package controllers

import (
	"encoding/json"
	"net/http"
	"taskmanager/common"
	"taskmanager/data"
	"taskmanager/models"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//POST /tasks
func CreateTask(w http.ResponseWriter, r *http.Request) {
	bodyData := getBodyFromRequest(w, r)
	task := &bodyData
	mc := NewMongoContext()
	defer mc.Close()
	context := mc.GetCollection("tasks")
	repo := &data.TaskRepository{context}
	err := repo.CreateTask(task)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Create task fail",
			http.StatusInternalServerError,
		)
	}
	if j, err := json.Marshal(TaskResource{Data: *task}); err != nil {
		common.DisplayUnexpectedAppError(w, err)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(j)
	}
}

//PUT /tasks/{id}
func UpdateTask(w http.ResponseWriter, r *http.Request) {
	//从url中获取id
	vars := mux.Vars(r)
	id := bson.ObjectIdHex(vars["id"])
	bodyData := getBodyFromRequest(w, r)
	task := &bodyData
	task.Id = id
	mc := NewMongoContext()
	defer mc.Close()
	context := mc.GetCollection("tasks")
	repo := &data.TaskRepository{context}
	if err := repo.UpdateTask(*task); err != nil {
		common.DisplayUnexpectedAppError(w, err)
	} else {
		w.WriteHeader(http.StatusNoContent)
	}

}

//GET /tasks
func GetTasks(w http.ResponseWriter, r *http.Request) {
	mc := NewMongoContext()
	defer mc.Close()
	context := mc.GetCollection("tasks")
	repo := &data.TaskRepository{context}
	tasks := repo.GetAllTasks()
	if j, err := json.Marshal(TasksResource{Data: tasks}); err != nil {
		common.DisplayUnexpectedAppError(w, err)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(j)
	}
}

//GET /tasks/{id}
func GetTaskById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	mc := NewMongoContext()
	context := mc.GetCollection("tasks")
	repo := &data.TaskRepository{context}
	task, err := repo.GetTaskById(id)
	if err != nil {
		if err == mgo.ErrNotFound {
			w.WriteHeader(http.StatusNoContent)
			return
		} else {
			common.DisplayUnexpectedAppError(w, err)
			return
		}
	}
	if j, err := json.Marshal(TaskResource{Data: task}); err != nil {
		common.DisplayUnexpectedAppError(w, err)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(j)
	}
}

//GET /tasks/users/{id}
func GetTaskByUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	mc := NewMongoContext()
	context := mc.GetCollection("tasks")
	repo := &data.TaskRepository{context}
	tasks := repo.GetTaskByUser(id)
	if j, err := json.Marshal(TasksResource{Data: tasks}); err != nil {
		common.DisplayUnexpectedAppError(w, err)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(j)
	}
}

//DELETE /tasks/{id}
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	mc := NewMongoContext()
	context := mc.GetCollection("tasks")
	repo := &data.TaskRepository{context}
	err := repo.DeleteTask(id)
	if err != nil {
		common.DisplayUnexpectedAppError(w, err)
	} else {
		w.WriteHeader(http.StatusNoContent)
	}
}

func getBodyFromRequest(w http.ResponseWriter, r *http.Request) models.Task {
	var taskResource TaskResource
	err := json.NewDecoder(r.Body).Decode(&taskResource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid task data",
			http.StatusInternalServerError,
		)
	}
	return taskResource.Data
}
