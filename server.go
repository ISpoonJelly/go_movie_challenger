package main

import (
	gin "github.com/gin-gonic/gin"

	"github.com/gin-contrib/sessions"
)

func main() {
	router := gin.Default()

	InitAPI(router)
	session := InitDB(router)
	defer session.Close()

	sessionColl := session.DB("movies").C("sessions")
	store := sessions.NewMongoStore(sessionColl, 3600, true, []byte(""))
	router.Use(sessions.Sessions("mysession", store))

	InitRoutes(router)

	router.Run(":4500")
}
