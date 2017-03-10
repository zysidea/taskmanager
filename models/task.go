package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Task struct {
	Id          bson.ObjectId `bson:"_id,omitempty" json:"id"`
	CreateBy    string        `json:"createby"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	CreateOn    time.Time     `json:"createon,omitempty"`
	Due         time.Time     `json:"due,omitempty"`
	Status      string        `json:"status,omitempty"`
	Tags        []string      `json:"tags,omitempty"`
}
