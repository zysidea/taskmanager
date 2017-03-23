package models

import (
	"gopkg.in/mgo.v2/bson"
)

type Task struct {
	Id          bson.ObjectId `bson:"_id,omitempty" json:"id"`
	CreateBy    string        `json:"createby"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	CreateOn    int64         `json:"createon,omitempty"`
	Due         int64         `json:"due,omitempty"`
	Status      string        `json:"status,omitempty"`
	Tags        []string      `json:"tags,omitempty"`
}
