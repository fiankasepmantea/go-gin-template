package main

import (
	"github.com/gin-gonic/gin"
	"github.com/fiankasepman/go-gin-template/internal/database"
)

func main() {

	database.Connect()
	database.RunMigrations()

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "server running",
		})
	})

	r.Run(":8080")
}