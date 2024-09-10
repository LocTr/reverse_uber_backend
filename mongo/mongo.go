package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

/// Study note: I guess this is to abtract the mongo connection. so other package can use this instead of mongo driver.

type Database interface {
	Collection(string) mongo.Collection
	Client() Client
}

type Client interface {
	Database(string) Database
	Connect(context.Context) error
	Disconnect(context.Context) error
}

func NewClient() Client {
	return &client{}
}
