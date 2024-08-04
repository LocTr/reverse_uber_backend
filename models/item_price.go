package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ItemPrice struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	ListPrice    float32            `bson:"list_price,omitempty"`
	CurrentPrice float32            `bson:"current_price,omitempty"`
	Description  string             `bson:"description,omitempty"`
}
