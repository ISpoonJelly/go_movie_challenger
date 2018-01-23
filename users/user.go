package users

import "gopkg.in/mgo.v2/bson"

type User struct {
	ID           bson.ObjectId `json:"-" bson:"_id,omitempty"`
	Username     string        `json:"username" bson:"username" binding:"required"`
	Name         string        `json:"name" bson:"name" binding:"required"`
	Password     string        `json:"password,omitempty" binding:"required"`
	PasswordHash string        `json:"-" bson:"hash"`
}

type LoginUser struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
