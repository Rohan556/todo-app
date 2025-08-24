package helper

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func GetUserIDFromContext(ctx *gin.Context) (bson.ObjectID, error) {
	userIDStr := ctx.GetString("userId")
	fmt.Println("UserID from context:", userIDStr)
	return bson.ObjectIDFromHex(userIDStr)
}
