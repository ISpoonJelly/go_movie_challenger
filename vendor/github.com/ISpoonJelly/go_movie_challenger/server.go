package main

import (
	gin "github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	session := InitDB(router)
	defer session.Close()

	InitRoutes(router)

	router.Run(":8080")
}
