package users

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func Init(router *gin.Engine) {
	router.GET("/user", getUsers)
	router.GET("/user/:username", getUser)
	router.POST("/user", createUser)
	router.POST("/login", loginUser)
}

func createUser(c *gin.Context) {
	dbColl := c.MustGet("DB").(*mgo.Database).C("users")

	var user User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, bson.M{"message": "Invalid parameters"})
		return
	}

	if err := dbColl.Find(bson.M{"username": user.Username}).One(nil); err == nil {
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

	err = dbColl.Insert(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, user)
}

func loginUser(c *gin.Context) {
	dbColl := c.MustGet("DB").(*mgo.Database).C("users")

	var login LoginUser

	if err := c.ShouldBind(&login); err != nil {
		c.JSON(http.StatusBadRequest, bson.M{"message": "Invalid parameters"})
		return
	}

	var user User
	if err := dbColl.Find(bson.M{"username": login.Username}).One(&user); err != nil {
		c.JSON(http.StatusUnauthorized, bson.M{"message": "User not found"})
		return
	}

	if !checkPasswordHash(login.Password, user.PasswordHash) {
		c.JSON(http.StatusUnauthorized, bson.M{"message": "Wrong Credentials"})
		return
	}

	session := sessions.Default(c)
	session.Set("user", user.ID.String())
	session.Save()

	c.JSON(http.StatusOK, "")
}

// func loggedIn(c *gin.Context) {
// 	session := sessions.Default(c)
// 	user := session.Get("user")

// 	c.JSON(http.StatusOK, user)
// }

func getUser(c *gin.Context) {
	dbColl := c.MustGet("DB").(*mgo.Database).C("users")

	username := c.Param("username")
	var result User
	err := dbColl.Find(bson.M{"username": username}).One(&result)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, result)
}

func getUsers(c *gin.Context) {
	dbColl := c.MustGet("DB").(*mgo.Database).C("users")

	var result []User
	err := dbColl.Find(nil).All(&result)

	if err != nil {
		c.JSON(http.StatusInternalServerError, bson.M{"error": err})
		return
	}

	c.JSON(http.StatusOK, result)
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func checkPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
