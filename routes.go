package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// InitRoutes initializes server routes
func InitRoutes(router *gin.Engine) {
	router.GET("/status", getStatus)
}

func getStatus(c *gin.Context) {
	c.String(http.StatusOK, "Server is up!")
}
