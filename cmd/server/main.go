package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/fiankasepman/go-gin-template/configs"
	"github.com/fiankasepman/go-gin-template/internal/database"
	"github.com/fiankasepman/go-gin-template/internal/middleware"
	userModule "github.com/fiankasepman/go-gin-template/internal/modules/user"
)

func main() {

	configs.LoadEnv()

	database.Connect()
	database.RunMigrations()
	database.SeedAll()

	var name string
	fmt.Println("DB NAME CHECK")
	database.DB.Raw("SELECT current_database()").Scan(&name)
	fmt.Println("CONNECTED DB:", name)

	r := gin.Default()

	db := database.DB

	userRepo := userModule.NewRepository(db)
	userService := userModule.NewService(userRepo)
	userHandler := userModule.NewHandler(userService)

	// PUBLIC
	r.POST("/login", userHandler.Login)

	// PROTECTED
	authGroup := r.Group("/")
	authGroup.Use(
		middleware.AuthMiddleware(),
		middleware.RBACMiddleware(db),
	)

	authGroup.GET("/users", userHandler.GetAll)
	authGroup.GET("/me", userHandler.Me)

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "server running"})
	})

	r.Run(":8080")
}