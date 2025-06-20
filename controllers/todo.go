package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rohan/go-todo/internal/schema"
)

func GetAllTodos() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"todos": []schema.Todo{
				{Id: 1, Completed: false, Title: "Learn Go"},
				{Id: 2, Completed: true, Title: "Build a web app"},
				{Id: 3, Completed: false, Title: "Write tests"},
			},
		})
	}
}
