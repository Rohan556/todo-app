package schema

type UserLoginSchema struct {
	Email    string `json:"email" bson:"email" validate:"required"`
	Password string `json:"password" bson:"password" validate:"required"`
}
