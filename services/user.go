package services

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/rohan/go-todo/database"
	"github.com/rohan/go-todo/internal/helper"
	"github.com/rohan/go-todo/internal/schema"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

const (
	UserCollectionName = "users"
)

func CreateUser(ctx *gin.Context, client *database.Database, document schema.CreateUserRequestBody) (*mongo.InsertOneResult, error) {
	collection := helper.ConnectToMongoDBCollection(client, Databasename, UserCollectionName)

	if document.Role == "" {
		document.Role = schema.User
	}

	return collection.InsertOne(ctx, document)
}

func GetUserInfo(ctx *gin.Context, client *database.Database, email string) (schema.UserSchema, error) {
	collection := helper.ConnectToMongoDBCollection(client, Databasename, UserCollectionName)

	fmt.Println(email, "email")

	var User schema.UserSchema

	err := collection.FindOne(ctx, bson.D{{"email", email}}).Decode(&User)

	return User, err
}
