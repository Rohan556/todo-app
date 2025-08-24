package schema

import "go.mongodb.org/mongo-driver/v2/bson"

type CreateToDoRequestBody struct {
	Completed bool          `bson:"completed" json:"completed"`
	Title     string        `bson:"title" json:"title" validate:"required"`
	UserID    bson.ObjectID `bson:"userId" json:"userId"`
}

type Todo struct {
	ID        bson.ObjectID `bson:"_id" json:"id"`
	Completed bool          `bson:"completed" json:"completed"`
	Title     string        `bson:"title" json:"title"`
}

type DeleteTodoRequestBody struct {
	ID []bson.ObjectID `json:"id"`
}
