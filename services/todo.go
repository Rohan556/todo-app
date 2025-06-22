package services

import (
	"github.com/gin-gonic/gin"
	"github.com/rohan/go-todo/database"
	"github.com/rohan/go-todo/internal/helper"
	"github.com/rohan/go-todo/internal/schema"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

const (
	Databasename   string = "todo"
	CollectionName string = "todo"
)

func GetTodos(ctx *gin.Context, client *database.Database) (*mongo.Cursor, error) {
	collection := helper.ConnectToMongoDBCollection(client, Databasename, CollectionName)

	return collection.Find(ctx, bson.M{})
}

func AddTodo(ctx *gin.Context, client *database.Database, requestBody schema.CreateToDoRequestBody) (*mongo.InsertOneResult, error) {
	collection := helper.ConnectToMongoDBCollection(client, Databasename, CollectionName)

	return collection.InsertOne(ctx, requestBody)
}
