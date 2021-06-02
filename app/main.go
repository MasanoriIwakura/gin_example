package main

import (
	"github.com/gin-gonic/gin"
	"app/controller"
)

func main() {
	router := gin.Default()
	
	router.GET("/health", controller.Health)

	router.Run(":8080")
}
