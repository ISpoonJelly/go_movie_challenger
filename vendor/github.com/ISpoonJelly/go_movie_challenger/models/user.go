package models

import "gopkg.in/mgo.v2/bson"

type User struct {
	ID       bson.ObjectId `bson:"_id,omitempty"`
	Username string        `json:"username" bson:"username" binding:"required"`
	Name     string        `json:"name" bson:"name" binding:"required"`
}
