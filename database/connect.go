package database

import (
	"context"
	"log"
	"todo-tasks/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect(databaseName string) *mongo.Database {
	uri := config.EnvMongoUri()
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	return client.Database(databaseName)
}

func GetCollection(collectionName string) *mongo.Collection {
	client := Connect("tasks")
	return client.Collection(collectionName)
}
