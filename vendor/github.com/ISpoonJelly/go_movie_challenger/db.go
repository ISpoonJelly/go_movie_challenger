package main

import (
	gin "github.com/gin-gonic/gin"
	mgo "gopkg.in/mgo.v2"
)

func InitDB(router *gin.Engine) *mgo.Session {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}

	db := session.DB("movies")

	router.Use(func(c *gin.Context) {
		c.Set("DB", db)
		c.Next()
	})

	return session
}
