package main

import "github.com/gin-gonic/gin"

import "net/http"

var port = ":8080"

func main() {
	router := gin.Default()

	router.GET("/status", getStatus)

	router.Run(port)
}

func getStatus(c *gin.Context) {
	c.String(http.StatusOK, "Server is up!")
}
