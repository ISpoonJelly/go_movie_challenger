package main

import (
	"net/http"

	controllers "github.com/ISpoonJelly/go_movie_challenger/controllers"

	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine) {
	router.GET("/status", getStatus)

	//User controllers
	router.GET("/user", controllers.GetUsers)
	router.GET("/user/:username", controllers.GetUser)
	router.POST("/user", controllers.CreateUser)
	router.POST("/login", controllers.LoginUser)
}

func getStatus(c *gin.Context) {
	c.String(http.StatusOK, "Server is up!")
}
