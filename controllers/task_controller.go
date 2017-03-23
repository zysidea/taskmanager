package controllers

import (
	"net/http"
	"encoding/json"
	"taskmanager/common"
	"taskmanager/data"
)

func CreateTask(w http.ResponseWriter,r *http.Request)  {
	var taskResource TaskResource
	err:=json.NewDecoder(r.Body).Decode(&taskResource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid task data",
			http.StatusInternalServerError,
		)
	}
	task:=&taskResource.Data
	mc:=NewContext()
	defer mc.Close()
	context:=mc.GetCollection("tasks")
	repo:=data.TaskRepository{context}
	err=repo.CreateTask(task)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Create task fail",
			http.StatusInternalServerError,
		)
	}
	if j,err:=json.Marshal(TaskResource{Data:*task});err!=nil{
		common.DisplayUnexceptAppError(w,err)
	}else {
		w.Header().Set("Content-Type","application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(j)
	}
}

func UpdateTask(w http.ResponseWriter,r *http.Request)  {

}
func GetTasks(w http.ResponseWriter,r *http.Request)  {

}
func GetTaskById(w http.ResponseWriter,r *http.Request)  {

}
func GetTaskByUser(w http.ResponseWriter,r *http.Request)  {

}
func DeleteTask(w http.ResponseWriter,r *http.Request)  {

}
