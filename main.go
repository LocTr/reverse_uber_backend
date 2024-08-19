package main

import (
	"context"
	"log"
	"net/url"

	"github.com/LocTr/reverse_uber_be/driver"
	"github.com/LocTr/reverse_uber_be/modules/config"
	"github.com/LocTr/reverse_uber_be/services"
	"github.com/gin-gonic/gin"
)

var app config.AppTools

func main() {

	services.LoadConfig()

	// Load environment configuration
	env, err := config.NewEnvConfig()
	if err != nil {
		log.Fatal("Failed to load environment configuration:", err)
	}
	// Construct the MongoDB connection URI
	uri := "mongodb+srv://" + url.QueryEscape(env.User) + ":" + url.QueryEscape(env.Password) +
		"@" + env.Cluster + "/?appName=" + env.AppName
	// Connect to the MongoDB deployment
	client := driver.Connect(uri)

	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			app.ErrorLogger.Fatal(err)
			return
		}
	}()

	// Create a new Gin router
	router := gin.New()

	// Define a route for the root URL
	router.GET("/", func(c *gin.Context) {
		c.String(200, "Hello hehe")
	})

	err = router.Run(":8080")
	if err != nil {
		app.ErrorLogger.Fatal(err)
	}
}
