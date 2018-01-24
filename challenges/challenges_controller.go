package challenges

import (
  "strconv"
  "fmt"

  mgo "gopkg.in/mgo.v2"
	"net/http"
	"github.com/dghubble/sling"
	"github.com/gin-gonic/gin"
  "github.com/ISpoonJelly/go_movie_challenger/movies"
	"gopkg.in/mgo.v2/bson"
  "github.com/gin-contrib/sessions"
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

  sorting_query_string := "&sort_by=vote_average.desc"
  voting_count_query_string := "&vote_count.gte=1000"
  req_params := genres_query_string + sorting_query_string + voting_count_query_string

  // check if there is an existing challenge for such user
  challenge := new(Challenge)
  tmdbErr := new(movies.TmdbError)
  _, err := api.Get("discover/movie?" + req_params).QueryStruct(apiParams).Receive(challenge, tmdbErr)

  if err != nil {
		c.JSON(http.StatusInternalServerError, bson.M{"error": err})
	} else {
    insertChallengeDB(c, challenge)
    c.JSON(http.StatusOK, challenge)
  }
}


func insertChallengeDB(c *gin.Context, challenge *Challenge) {
  session := sessions.Default(c)
  id := session.Get("user").(string)

  challenge.User = id
  challengeCollDB := c.MustGet("DB").(*mgo.Database).C("challenges")
  err := challengeCollDB.Insert(challenge)
  if err != nil {
    fmt.Println("ERR --> ", err)
    c.JSON(http.StatusInternalServerError, err)
		return
  }
}
