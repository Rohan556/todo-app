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
	"github.com/rohan/go-todo/internal/routes"
)

type Book struct {
	Title  string
	Author string
}

func main() {
	fmt.Println("Loading env variables...")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT := os.Getenv("PORT")

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

	server.GET(routes.TODO_URL, controllers.GetAllTodos(&databaseClient))
	server.POST(routes.TODO_URL, controllers.AddTodo(&databaseClient))
	server.POST(routes.USER_URL, controllers.CreateUser(&databaseClient))
	server.POST(routes.LOGIN_URL, controllers.LoginUser(&databaseClient))

	server.Run(":" + PORT)
}
