package controllers

import (
	"net/http"

	"github.com/ISpoonJelly/go_movie_challenger/models"

	"github.com/gin-gonic/gin"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func GetPeople(c *gin.Context) {
	db, ok := c.Keys["DB"].(*mgo.Database)

	if !ok {
		panic("db not found")
	}

	var result []models.Person
	err := db.C("people").Find(bson.M{}).All(&result)

	if err != nil {
		c.JSON(http.StatusInternalServerError, bson.M{"error": err})
		return
	}

	c.JSON(http.StatusOK, bson.M{"people": result})
}
