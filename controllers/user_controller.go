package controllers

import (
	"net/http"

	"github.com/ISpoonJelly/go_movie_challenger/models"
	"golang.org/x/crypto/bcrypt"

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

	if err := coll.Find(bson.M{"username": user.Username}).One(nil); err == nil {
		c.JSON(http.StatusBadRequest, bson.M{"message": "User already registered"})
		return
	}

	hashed, err := hashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	user.PasswordHash = hashed
	user.Password = ""

	err = coll.Insert(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, user)
}

func LoginUser(c *gin.Context) {
	db, ok := c.Keys["DB"].(*mgo.Database)
	if !ok {
		panic("db not found")
	}

	var login models.LoginUser

	if err := c.ShouldBind(&login); err != nil {
		c.JSON(http.StatusBadRequest, bson.M{"message": "Invalid parameters"})
		return
	}

	var user models.User
	if err := db.C("user").Find(bson.M{"username": login.Username}).One(&user); err != nil {
		c.JSON(http.StatusUnauthorized, bson.M{"message": "User not found"})
		return
	}

	if !checkPasswordHash(login.Password, user.PasswordHash) {
		c.JSON(http.StatusUnauthorized, bson.M{"message": "Wrong Credentials"})
		return
	}

	//m3ana el user
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

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
