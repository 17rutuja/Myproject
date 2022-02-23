package main

import (
	"golang_crudapp/routes"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5001"
	}

	router := gin.New()

	// for logging purpose
	router.Use(gin.Logger())

	// routes
	routes.Routes(router)

	// listening server
	router.Run(":" + port)
}
