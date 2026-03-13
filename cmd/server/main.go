package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
	"github.com/fiankasepman/go-gin-template/internal/database"
)

func main() {

	godotenv.Load()

	err := database.Connect()
	if err != nil {
		panic(err)
	}

	r := gin.Default()

	port := os.Getenv("APP_PORT")

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "server running",
		})
	})

	fmt.Println("server running :" + port)

	r.Run(":" + port)
}