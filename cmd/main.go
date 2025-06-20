package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rohan/go-todo/controllers"
)

func main() {
	fmt.Println("Loading env variables...")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT := os.Getenv("PORT")

	fmt.Println("Env variables loaded successfully")
	fmt.Println("Starting the server...")
	server := gin.Default()
	server.GET("/", controllers.GetAllTodos())

	server.Run(":" + PORT)
}
