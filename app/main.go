package main

import (
	"app/controllers"
	"app/database"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	db, _ := database.Initialize()

	router := gin.Default()
	router.Use(database.Inject(db))

	router.GET("/health", controllers.Health)
	router.GET("/sample", controllers.SampleIndex)

	router.Run(":8080")
}
