package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load environment variables
	port := os.Getenv("WEB_PORT")
	if port == "" {
		port = "3140" // Default port
	}

	// Initialize Gin web server
	router := gin.Default()

	// Serve the main page
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to Re:Delivery!",
		})
	})

	// Start the web server
	log.Printf("Re:Delivery running on port %s", port)
	err := router.Run("0.0.0.0:" + port)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
