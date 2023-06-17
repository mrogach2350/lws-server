package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/mrogach2350/lws-server/handlers"
	"github.com/mrogach2350/lws-server/models"
)

func main() {
	godotenv.Load(".env")

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	tempSecret := os.Getenv("TEMP_SECRET")
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/defaultdb?sslmode=verify-full", dbUser, dbPassword, dbHost, dbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("failed to connect database", err)
	}

	db.AutoMigrate(&models.List{}, &models.Movie{}, &models.Genre{})

	r := gin.Default()
	r.Use(cors.Default())
	r.Use(func(c *gin.Context) {
		bearerToken := c.Request.Header.Get("Authorization")
		if len(strings.Split(bearerToken, " ")) == 2 {
			token := strings.Split(bearerToken, " ")[1]
			if token == tempSecret {
				c.Next()
				return
			}
		}
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "Invalid Bearer Token!"})
	})

	r.GET("/movies", handlers.GetAllMovies(db))
	r.POST("/movies", handlers.CreateNewMovie(db))
	r.GET("/movies/:id", handlers.GetMovieById(db))
	r.POST("/movies/:id/watched", handlers.SetMovieWatched(db))
	r.GET("/lists", handlers.GetAllLists(db))
	r.GET("/lists/:listId/movies", handlers.GetMoviesByListId(db))
	r.POST("/lists", handlers.CreateNewList(db))

	r.Run()
}
