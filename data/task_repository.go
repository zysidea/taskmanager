package data

import (
	"gopkg.in/mgo.v2"
	"taskmanager/models"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type TaskRepository struct {
	C *mgo.Collection
}

func (r *TaskRepository) CreateTask(task *models.Task) error {
	objId := bson.NewObjectId()
	task.Id = objId
	task.CreateOn = time.Now()
	task.Status = "Created"
	err := r.C.Insert(task)
	return err
}
func (r *TaskRepository) UpdateTask(task models.Task) error {
	taskId := task.Id
	err := r.C.Update(bson.M{"_id": taskId}, bson.M{
		"$set": bson.M{
			"name":        task.Name,
			"description": task.Description,
			"due":         task.Due,
			"status":      task.Status,
			"tags":        task.Tags,
		},
	})
	return err
}
func (r *TaskRepository) DeleteTask(id string) error {
	err := r.C.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
}

func (r *TaskRepository) GetAllTasks() []models.Task {
	var tasks []models.Task
	iter := r.C.Find(nil).Iter()
	result := models.Task{}
	for iter.Next(&result) {
		tasks = append(tasks, result)
	}
	return tasks
}
func (r *TaskRepository) GetTaskById(id string)(task models.Task,err error){
	err= r.C.Find(bson.M{"_id":bson.ObjectIdHex(id)}).One(&task)
	return
}
func (r *TaskRepository) GetTaskByUser(user string) []models.Task  {
	var tasks []models.Task
	iter:=r.C.Find(bson.M{"createby":user}).Iter()
	result:=models.Task{}
	for iter.Next(result) {
		tasks=append(tasks,result)
	}
	return tasks
}