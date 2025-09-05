package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rohan/go-todo/controllers"
	"github.com/rohan/go-todo/database"
	"github.com/rohan/go-todo/internal/middlewares"
	"github.com/rohan/go-todo/internal/routes"
)

func main() {
	fmt.Println("Loading env variables...")

	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}

	fmt.Println("Env variables loaded successfully")

	MONGO_URI := os.Getenv("MONGO_URI")

	databaseClient := database.ConnectToMongoDB(MONGO_URI)

	defer func() {
		if err = databaseClient.Client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	if databaseClient.Client == nil {
		log.Fatal("Failed to connect to MongoDB")
	}

	fmt.Println("Starting the server...")
	server := gin.Default()

	server.POST(routes.USER_URL, controllers.CreateUser(&databaseClient))
	server.POST(routes.LOGIN_URL, controllers.LoginUser(&databaseClient))

	protected := server.Group("/")
	protected.Use(middlewares.JWTAuthMiddlewares())
	{
		protected.GET(routes.TODO_URL, controllers.GetAllTodos(&databaseClient))
		protected.POST(routes.TODO_URL, controllers.AddTodo(&databaseClient))
		protected.DELETE(routes.TODO_URL, controllers.DeleteTodo(&databaseClient))
		protected.PUT(routes.TODO_URL, controllers.UpdateTodo(&databaseClient))
	}

	server.Run(":" + PORT)
}
