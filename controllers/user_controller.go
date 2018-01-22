package controllers

import (
	"net/http"

	"github.com/ISpoonJelly/go_movie_challenger/models"

	"github.com/gin-gonic/gin"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func CreateUser(c *gin.Context) {
	db, ok := c.Keys["DB"].(*mgo.Database)

	if !ok {
		panic("db not found")
	}

	var user models.User

	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, bson.M{"message": "Invalid parameters"})
		return
	}

	coll := db.C("user")

	if err := db.C("user").Find(bson.M{"username": user.Username}).One(nil); err == nil {
		c.JSON(http.StatusBadRequest, bson.M{"message": "User already registered"})
		return
	}

	err := coll.Insert(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, bson.M{"e2": err})
		return
	}

	c.JSON(http.StatusOK, user)
	return
}

func GetUser(c *gin.Context) {
	db, ok := c.Keys["DB"].(*mgo.Database)

	if !ok {
		panic("db not found")
	}

	username := c.Param("username")
	var result models.User
	err := db.C("user").Find(bson.M{"username": username}).One(&result)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, result)
}

func GetUsers(c *gin.Context) {
	db, ok := c.Keys["DB"].(*mgo.Database)

	if !ok {
		panic("db not found")
	}

	var result []models.User
	err := db.C("user").Find(nil).All(&result)

	if err != nil {
		c.JSON(http.StatusInternalServerError, bson.M{"error": err})
		return
	}

	c.JSON(http.StatusOK, result)
}
