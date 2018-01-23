package movies

import (
	"net/http"

	"github.com/dghubble/sling"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

func Init(router *gin.Engine) {
	router.GET("/genres", getGenres)
}

func getGenres(c *gin.Context) {
	api := c.MustGet("tmdbAPI").(*sling.Sling)
	key := c.MustGet("tmdbKey").(string)

	apiParams := APIParams{APIKey: key}

	genres := new(Genres)
	tmdbErr := new(TmdbError)
	_, err := api.Get("genre/movie/list").QueryStruct(apiParams).Receive(genres, tmdbErr)

	if err != nil {
		c.JSON(http.StatusInternalServerError, bson.M{"error": err})
	}

	c.JSON(http.StatusOK, genres)
}
