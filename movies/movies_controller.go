package movies

import (
	"fmt"

	"github.com/dghubble/sling"
	"github.com/gin-gonic/gin"
)

func Init(router *gin.Engine) {
	router.GET("/genres", getGenres)
}

func getGenres(c *gin.Context) {
	api := c.MustGet("tmdbAPI").(*sling.Sling)
	key := c.MustGet("tmdbKey").(string)

	apiParams := APIParams{APIKey: key}

	genres := new(Genre)
	tmdbErr := new(TmdbError)
	resp, err := api.Get("genre/movie/list").QueryStruct(apiParams).Receive(&genres, &tmdbErr)

	fmt.Println("REQUEST: ", resp.Request, "\n\nRESPONSE: ", resp, "\n\nERR: ", err, "\n\nGENRES: ", genres, "\n\nTMDBERR: ", tmdbErr)

	if err != nil {
		panic(err)
	}
}
