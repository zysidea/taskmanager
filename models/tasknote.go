package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type TaskNote struct {
	Id          bson.ObjectId `bson:"_id,omitempty" json:"id"`
	TaskId      bson.ObjectId `json:"taskid"`
	Description string        `json:"description"`
	CreateOn    time.Time     `json:"createon,omitempty"`
}
