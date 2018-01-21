package main

import (
	gin "github.com/gin-gonic/gin"
	mgo "gopkg.in/mgo.v2"
)

func main() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	router := gin.Default()
	InitRoutes(router)

	router.Run(":8080")
}
