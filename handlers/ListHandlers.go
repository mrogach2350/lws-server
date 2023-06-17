package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mrogach2350/lws-server/models"
	"gorm.io/gorm"
)

func GetAllLists(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var lists []models.List
		db.Find(&lists)

		c.JSON(http.StatusOK, gin.H{"success": true, "data": lists})
	}
}

func GetMoviesByListId(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		listId := c.Param("listId")
		listIdUint, err := strconv.ParseUint(listId, 10, 64)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"success": false, "message": err.Error()})
		}

		var movies []models.Movie
		db.Where(&models.Movie{ListID: uint(listIdUint)}).Find(&movies)
		c.JSON(http.StatusOK, gin.H{"success": true, "data": movies})
	}
}

func CreateNewList(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var newList models.List

		c.BindJSON(&newList)
		db.Create(&newList)

		c.JSON(http.StatusOK, gin.H{"success": true, "data": newList})
	}
}
