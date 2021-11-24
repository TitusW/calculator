package main

import (
	"golang-calculator/routes"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "5000"
	}

	router := gin.New()
	router.Use(gin.Logger())
	routes.CalculatorRoutes(router)

	router.Run(":" + port)
}
