package main

import (
	"net/http"

	"github.com/ISpoonJelly/go_movie_challenger/movies"
	"github.com/ISpoonJelly/go_movie_challenger/users"
	"github.com/ISpoonJelly/go_movie_challenger/challenges"
	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine) {
	router.GET("/status", getStatus)

	users.Init(router)
	movies.Init(router)
	challenges.Init(router)
}

func getStatus(c *gin.Context) {
	c.String(http.StatusOK, "Server is up!")
}
