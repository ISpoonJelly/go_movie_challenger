package challenges

import (
  "fmt"
	"net/http"
	"github.com/dghubble/sling"
	"github.com/gin-gonic/gin"
  "github.com/ISpoonJelly/go_movie_challenger/movies"
	"gopkg.in/mgo.v2/bson"
)

func Init(router *gin.Engine) {
  	router.GET("/challenge", getChallenge)
}

func getChallenge(c *gin.Context) {
  api := c.MustGet("tmdbAPI").(*sling.Sling)
	key := c.MustGet("tmdbKey").(string)

  apiParams := movies.APIParams{APIKey: key}

  challenge := new(Challenge)
  tmdbErr := new(movies.TmdbError)
  res, err := api.Get("discover/movie").QueryStruct(apiParams).Receive(challenge, tmdbErr)

  if err != nil {
		c.JSON(http.StatusInternalServerError, bson.M{"error": err})
	} else {
    fmt.Println("RECEIVED --> ", challenge)
    fmt.Println("RES --> ", res)
  }
}
