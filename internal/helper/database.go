package helper

import (
	"github.com/rohan/go-todo/database"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func ConnectToMongoDBCollection(client *database.Database, databaseName, collectionName string) *mongo.Collection {
	collection := client.Client.Database(databaseName).Collection(collectionName)

	return collection
}
