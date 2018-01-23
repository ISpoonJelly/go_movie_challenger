package challenges

import (
  "strconv"
  "fmt"
	"net/http"
	"github.com/dghubble/sling"
	"github.com/gin-gonic/gin"
  "github.com/ISpoonJelly/go_movie_challenger/movies"
	"gopkg.in/mgo.v2/bson"
)

func Init(router *gin.Engine) {
  	router.POST("/challenge", getChallenge)
}

func getChallenge(c *gin.Context) {
  api := c.MustGet("tmdbAPI").(*sling.Sling)
	key := c.MustGet("tmdbKey").(string)

  apiParams := movies.APIParams{APIKey: key}

  genre_ids := new(Genre_ids)
  if err := c.ShouldBind(&genre_ids); err != nil {
		c.JSON(http.StatusBadRequest, bson.M{"message": "Invalid parameters"})
		return
	}

  var genres_query_string string
  for i := 0; i < len(genre_ids.Genre_ids); i++ {
    if i == 0 {
      genres_query_string = "&with_genres=" + strconv.Itoa(genre_ids.Genre_ids[i])
    } else {
        genres_query_string += "|with_genres=" + strconv.Itoa(genre_ids.Genre_ids[i])
    }
  }

  // DEBUGGING --> FORMATTING QUERY STRING FOR WITH_GENRES
  fmt.Println("QUERY GENRES --> ", genres_query_string)


  challenge := new(Challenge)
  tmdbErr := new(movies.TmdbError)
  _, err := api.Get("discover/movie?" + genres_query_string).QueryStruct(apiParams).Receive(challenge, tmdbErr)

  if err != nil {
		c.JSON(http.StatusInternalServerError, bson.M{"error": err})
	} else {
    c.JSON(http.StatusOK, challenge)
  }
}
