package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type MenuItem struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name,omitempty"`
	Description string             `bson:"description,omitempty"`
	Price       float32            `bson:"price,omitempty"`
	Tags        []string           `bson:"tags,omitempty"`
}
