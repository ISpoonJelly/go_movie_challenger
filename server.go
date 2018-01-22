package main

import (
	gin "github.com/gin-gonic/gin"

	"github.com/gin-contrib/sessions"
)

func main() {
	router := gin.Default()

	session := InitDB(router)
	defer session.Close()

	sessionColl := session.DB("movies").C("sessions")
	store := sessions.NewMongoStore(sessionColl, 3600, true, []byte("secret"))
	router.Use(sessions.Sessions("user-session", store))

	InitRoutes(router)

	router.Run(":8080")
}
