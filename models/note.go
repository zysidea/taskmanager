package models

import (

	"gopkg.in/mgo.v2/bson"
)

type Note struct {
	Id          bson.ObjectId `bson:"_id,omitempty" json:"id"`
	TaskId      bson.ObjectId `json:"taskid"`
	Description string        `json:"description"`
	CreateOn    int64     `json:"createon,omitempty"`
}
