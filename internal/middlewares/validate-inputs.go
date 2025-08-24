package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/rohan/go-todo/internal/loggers"
)

func HandleInputValidations(requestBody interface{}) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.BindJSON(&requestBody)

		validationSuccess := loggers.ValidateRequestBody(ctx, requestBody)

		if !validationSuccess {
			return
		}
		ctx.Next()
	}
}
