package models

import "gopkg.in/mgo.v2/bson"

type User struct {
	ID           bson.ObjectId `json:"-" bson:"-"`
	Username     string        `json:"username" bson:"username" binding:"required"`
	Name         string        `json:"name" bson:"name" binding:"required"`
	Password     string        `json:"password,omitempty" bson:"password,omitempty" binding:"required"`
	PasswordHash string        `json:"-" bson:"hash"`
}
