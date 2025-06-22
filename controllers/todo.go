package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rohan/go-todo/database"
	"github.com/rohan/go-todo/internal/loggers"
	"github.com/rohan/go-todo/internal/schema"
	"github.com/rohan/go-todo/services"
)

type Book struct {
	Title  string
	Author string
}

func GetAllTodos(client *database.Database) gin.HandlerFunc {
	return func(c *gin.Context) {

		cursor, err := services.GetTodos(c, client)

		if err != nil {
			loggers.HandleResponse(c, http.StatusInternalServerError, err)
		}

		var todos []schema.Todo

		if err := cursor.All(c, &todos); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode todos"})
			return
		}

		loggers.HandleResponse(c, http.StatusOK, todos)
	}
}

func AddTodo(client *database.Database) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var requestBody schema.CreateToDoRequestBody

		ctx.BindJSON(&requestBody)

		validationSuccess := loggers.ValidateRequestBody(ctx, requestBody)

		if !validationSuccess {
			return
		}

		result, err := services.AddTodo(ctx, client, requestBody)

		if err != nil {
			loggers.HandleResponse(ctx, http.StatusInternalServerError, err)
		}

		loggers.HandleResponse(ctx, http.StatusCreated, result)
	}
}
