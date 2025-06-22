package schema

import "go.mongodb.org/mongo-driver/v2/bson"

type Roles string

const (
	User  Roles = "user"
	Admin Roles = "admin"
)

type CreateUserRequestBody struct {
	Email    string `json:"email" bson:"email" validate:"required"`
	Password string `json:"password" bson:"password" validate:"required"`
	Role     Roles  `json:"role" bson:"role"`
	Name     string `json:"name" bson:"name" validate:"required"`
}

type UserSchema struct {
	ID    bson.ObjectID `bson:"_id" json:"id"`
	Email string        `json:"email" bson:"email"`
	Role  Roles         `json:"role" bson:"role"`
	Name  string        `json:"name" bson:"name"`
}
