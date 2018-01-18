package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Person struct {
	Name  string
	Phone string
}

var collPeople *mgo.Collection

func main() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	collPeople = session.DB("test").C("people")
	err = collPeople.Insert(&Person{"Tarek", "+201111984394"},
		&Person{"3amooor", "+201223340190"})

	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()

	router.GET("/status", getStatus)
	router.GET("/people", getPeople)

	router.Run(":8080")
}

func getStatus(c *gin.Context) {
	c.String(http.StatusOK, "Server is up!")
}

func getPeople(cntx *gin.Context) {
	var result []Person
	err := collPeople.Find(bson.M{}).All(&result)

	if err != nil {
		cntx.JSON(http.StatusInternalServerError, bson.M{"error": err})
		return
	}

	cntx.JSON(http.StatusOK, bson.M{"people": result})
}
