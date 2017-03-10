package models

import "gopkg.in/mgo.v2/bson"

type User struct {
	Id           bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Name         string        `json:"name"`
	Email        string        `json:"email"`
	Password     string        `json:"password,omitempty"`
	HashPassword []byte        `json:"hashpassword,omitempty"`
}
