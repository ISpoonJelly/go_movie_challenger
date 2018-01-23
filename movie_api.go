package main

import (
	"net/http"

	"github.com/dghubble/sling"
	gin "github.com/gin-gonic/gin"
)

const tmdbAPI = "https://api.themoviedb.org/3/"

func InitAPI(router *gin.Engine) {
	baseClient := sling.New().Base(tmdbAPI).Client(http.DefaultClient)

	router.Use(func(c *gin.Context) {
		c.Set("tmdbKey", "9e441b31135ec60b8e9fb96c24cdc0a6")
		c.Set("tmdbAPI", baseClient)
		c.Next()
	})
}
