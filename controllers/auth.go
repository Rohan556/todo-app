package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rohan/go-todo/database"
	"github.com/rohan/go-todo/internal/auth"
	"github.com/rohan/go-todo/internal/loggers"
	"github.com/rohan/go-todo/internal/schema"
	"github.com/rohan/go-todo/services"
)

func LoginUser(client *database.Database) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var requestBody schema.UserLoginSchema

		ctx.BindJSON(&requestBody)

		validationSuccess := loggers.ValidateRequestBody(ctx, requestBody)

		if !validationSuccess {
			return
		}

		user, err := services.GetUserInfo(ctx, client, requestBody.Email)

		isPasswordValid := auth.IsPasswordValid(user.Password, requestBody.Password)

		if !isPasswordValid {
			loggers.HandleResponse(ctx, http.StatusBadRequest, gin.H{
				"message": "User password is invalid",
			})
			return
		}

		if err == nil {
			payload := schema.JWTRequiredFields{
				Email:  user.Email,
				UserId: user.ID.Hex(),
			}

			tokenString, err := auth.GenerateJWTToken(payload)

			if err != nil {
				loggers.HandleResponse(ctx, http.StatusInternalServerError, err)
				return
			}

			loggers.HandleResponse(ctx, http.StatusOK, tokenString)
		} else {
			loggers.HandleResponse(ctx, http.StatusForbidden, gin.H{
				"message": "User does not exist",
			})
		}

	}
}
