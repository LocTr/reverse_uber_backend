package bootstrap

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoDatabase(env *Env) mongo.Client {

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(env.ContextTimeout)*time.Second)

	defer cancel()

	dbHost := env.DBHost
	dbPort := env.DBPort
	dbUser := env.DBUser
	dbPassword := env.DBPassword

	mongodbURI := fmt.Sprintf("mongodb://%s:%s@%s:%s", dbUser, dbPassword, dbHost, dbPort)

	clientOption := options.Client().ApplyURI(mongodbURI)

	client, err := mongo.Connect(ctx, clientOption)

	if err != nil {
		panic(err)
	}

	return client
}
