package query

import "go.mongodb.org/mongo-driver/mongo"

func Menu(db mongo.Client, collection string) *mongo.Collection {
	var menu = db.Database("reverse_uber").Collection(collection)
	return menu
}
