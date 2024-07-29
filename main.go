package main

import (
	"context"
	"fmt"
	"net/url"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Create a new Gin router
	router := gin.Default()

	// Define a route for the root URL
	router.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, World hehe")
	})

	godotenv.Load()
	cluster := os.Getenv("CLUSTER_NAME")
	appName := os.Getenv("APP_NAME")
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")

	uri := "mongodb+srv://" + url.QueryEscape(user) + ":" + url.QueryEscape(password) +
		"@" + cluster + "/?appName=" + appName

	println(uri)
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	// Create a new client and connect to the server
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	// Send a ping to confirm a successful connection
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	router.Run(":8080")
}
