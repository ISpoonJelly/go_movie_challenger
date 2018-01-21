package main

import (
	"net/http"

	controllers "github.com/ISpoonJelly/go_movie_challenger/controllers"

	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine) {
	router.GET("/status", getStatus)
	router.GET("/people", controllers.GetPeople)
}

func getStatus(c *gin.Context) {
	c.String(http.StatusOK, "Server is up!")
}
