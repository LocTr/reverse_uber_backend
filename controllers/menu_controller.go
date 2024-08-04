package controllers

import "go.mongodb.org/mongo-driver/mongo"

var menuCollection *mongo.Collection = configs.GetCollection(configs.DB, "menu")
