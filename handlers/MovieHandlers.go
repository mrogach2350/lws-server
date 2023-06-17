package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mrogach2350/lws-server/models"
	"gorm.io/gorm"
)

func GetAllMovies(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var movies []models.Movie
		db.Find(&movies)

		c.JSON(http.StatusOK, gin.H{"success": true, "data": movies})
	}
}

type NewMovieRequest struct {
	Movie     models.Movie `json:"movie"`
	ListValue string       `json:"list"`
}

func CreateNewMovie(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var newMovieRequest NewMovieRequest
		var list models.List

		c.BindJSON(&newMovieRequest)
		db.Where("value = ?", newMovieRequest.ListValue).First(&list)
		newMovie := newMovieRequest.Movie
		newMovie.ListID = list.ID

		db.Create(&newMovie)

		c.JSON(http.StatusOK, gin.H{"success": true, "data": newMovie})
	}
}

func GetMovieById(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		movieId := c.Param("id")
		var movie models.Movie

		db.Where("id = ?", movieId).First(&movie)

		c.JSON(http.StatusOK, gin.H{"success": true, "data": movie})
	}
}

type WatchedRequest struct {
	Watched bool `json:"watched"`
}

func SetMovieWatched(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		movieId := c.Param("id")
		var watchRequest WatchedRequest
		var movieRecord models.Movie
		c.BindJSON(&watchRequest)
		db.Where("id = ?", movieId).First(&movieRecord)
		movieRecord.Watched = watchRequest.Watched
		db.Save(&movieRecord)
		c.JSON(http.StatusOK, gin.H{"success": true, "data": movieRecord})
	}
}
