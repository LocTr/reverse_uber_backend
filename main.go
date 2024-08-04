package main

import (
	"fmt"
	"log"
	"net/url"

	"github.com/LocTr/reverse_uber_be/modules/config"
	"github.com/gin-gonic/gin"
)

func main() {
	// Create a new Gin router
	router := gin.Default()

	// Define a route for the root URL
	router.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, World hehe")
	})

	// Load environment configuration
	env, err := config.NewEnvConfig()
	if err != nil {
		log.Fatal("Failed to load environment configuration:", err)
	}
	// Construct the MongoDB connection URI
	uri := "mongodb+srv://" + url.QueryEscape(env.User) + ":" + url.QueryEscape(env.Password) +
		"@" + env.Cluster + "/?appName=" + env.AppName

	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	router.Run(":8080")
}
