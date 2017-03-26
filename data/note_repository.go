package data

import (
	"gopkg.in/mgo.v2"
	"taskmanager/models"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type NoteRepository struct {
	C *mgo.Collection
}

func (r *NoteRepository)CreateNote(note *models.Note) error {
	objId:=bson.NewObjectId()
	note.Id=objId
	note.CreateOn=time.Now().Unix()
	err:=r.C.Insert(note)
	return err
}

func (r *NoteRepository)UpdateNote(note *models.Note) error  {
	err:=r.C.Update(bson.M{"_id":note.Id},bson.M{
		"$set":bson.M{
			"description":note.Description,
		},
	})
	return err
}

func (r *NoteRepository) GetAllNotes() []models.Note {
	var notes []models.Note
	iter := r.C.Find(nil).Iter()
	result := models.Note{}
	for iter.Next(&result) {
		notes = append(notes, result)
	}
	return notes
}


func (r *NoteRepository)GetNoteById(id string) (note models.Note,err error)  {
	err = r.C.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&note)
	return
}

func (r *NoteRepository)GetNoteByTask(taskId string)[]models.Note  {
	var notes []models.Note
	iter := r.C.Find(bson.M{"createby": taskId}).Iter()
	result := models.Note{}
	for iter.Next(&result) {
		notes = append(notes, result)
	}
	return notes
}

func (r *NoteRepository) DeleteNote(id string) error {
	err := r.C.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
}
