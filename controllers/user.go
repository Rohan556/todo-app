package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rohan/go-todo/database"
	"github.com/rohan/go-todo/internal/auth"
	"github.com/rohan/go-todo/internal/loggers"
	"github.com/rohan/go-todo/internal/schema"
	"github.com/rohan/go-todo/services"
)

func CreateUser(client *database.Database) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var requestBody schema.CreateUserRequestBody
		ctx.BindJSON(&requestBody)

		validationSuccess := loggers.ValidateRequestBody(ctx, requestBody)

		if !validationSuccess {
			return
		}

		user, err := services.GetUserInfo(ctx, client, requestBody.Email)

		if err == nil {
			loggers.HandleResponse(ctx, http.StatusBadRequest, "User already exists")
			return
		}

		fmt.Println(user, "user")

		encryptedPassword, err := auth.EncryptPassword(requestBody.Password)

		if err != nil {
			loggers.HandleResponse(ctx, http.StatusInternalServerError, err)
			return
		}

		requestBody.Password = encryptedPassword

		_, err = services.CreateUser(ctx, client, requestBody)

		if err != nil {
			loggers.HandleResponse(ctx, http.StatusInternalServerError, err)
			return
		}

		loggers.HandleResponse(ctx, http.StatusCreated, "User created successfully")
	}
}
