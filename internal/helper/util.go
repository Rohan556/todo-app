package helper

import (
	"bytes"
	"encoding/json"
	"io"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func GetUserIDFromContext(ctx *gin.Context) (bson.ObjectID, error) {
	userIDStr := ctx.GetString("userId")
	return bson.ObjectIDFromHex(userIDStr)
}

// JSONHasKey checks if a JSON byte array contains a specific key
func JSONHasKey(data []byte, key string) bool {
	var m map[string]interface{}
	if err := json.Unmarshal(data, &m); err != nil {
		return false
	}
	_, exists := m[key]
	return exists
}

// CreateReadCloser creates a new io.ReadCloser from a byte slice
func CreateReadCloser(data []byte) io.ReadCloser {
	return io.NopCloser(bytes.NewBuffer(data))
}
