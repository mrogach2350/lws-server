package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/mrogach2350/lws-server/handlers"
	"github.com/mrogach2350/lws-server/models"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/defaultdb?sslmode=verify-full", dbUser, dbPassword, dbHost, dbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("failed to connect database", err)
	}

	db.AutoMigrate(&models.List{}, &models.Movie{}, &models.Genre{})

	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/movies", handlers.GetAllMovies(db))
	r.POST("/movies", handlers.CreateNewMovie(db))
	r.GET("/movies/:id", handlers.GetMovieById(db))
	r.POST("/movies/:id/watched", handlers.SetMovieWatched(db))
	r.GET("/lists", handlers.GetAllLists(db))
	r.GET("/lists/:listId/movies", handlers.GetMoviesByListId(db))
	r.POST("/lists", handlers.CreateNewList(db))

	r.Run()
}
