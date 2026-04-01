package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/fiankasepman/go-gin-template/configs"
	"github.com/fiankasepman/go-gin-template/internal/database"
	"github.com/fiankasepman/go-gin-template/internal/middleware"

	groupModule "github.com/fiankasepman/go-gin-template/internal/modules/group"
	userModule "github.com/fiankasepman/go-gin-template/internal/modules/user"
	usertoken "github.com/fiankasepman/go-gin-template/internal/modules/user_token"
)

func main() {

	// ================== LOAD ENV ==================
	configs.LoadEnv()

	// ================== DB ==================
	database.Connect()
	database.RunMigrations()
	database.SeedAll()

	// ================== CHECK DB ==================
	var name string
	fmt.Println("DB NAME CHECK")
	database.DB.Raw("SELECT current_database()").Scan(&name)
	fmt.Println("CONNECTED DB:", name)

	// ================== GIN ==================
	r := gin.Default()
	db := database.DB

	// ================== USER TOKEN MODULE ==================
	userTokenRepo := usertoken.NewRepository(db)

	// start cron
	usertoken.StartCleanupJob(userTokenRepo)

	// ================== USER MODULE ==================
	userRepo := userModule.NewRepository(db)
	userService := userModule.NewService(userRepo, userTokenRepo)
	userHandler := userModule.NewHandler(userService)

	// ================== GROUP MODULE ==================
	groupRepo := groupModule.NewRepository(db)
	groupService := groupModule.NewService(groupRepo)
	groupHandler := groupModule.NewHandler(groupService)

	// ================== PUBLIC ROUTES ==================
	r.POST("/login", userHandler.Login)
	r.POST("/refresh", userHandler.Refresh)

	// ================== PROTECTED ROUTES ==================
	authGroup := r.Group("/")
	authGroup.Use(
		// middleware.AuthMiddleware(),
		middleware.PasetoMiddleware(),
		middleware.RBACMiddleware(db),
	)

	// ---------- AUTH ----------
	authGroup.POST("/logout", userHandler.Logout)
	authGroup.POST("/logout-all", userHandler.LogoutAll)

	// ---------- USER ----------
	authGroup.GET("/users", userHandler.GetAll)
	authGroup.POST("/users", userHandler.Create)
	authGroup.PUT("/users/:id", userHandler.Update)
	authGroup.DELETE("/users/:id", userHandler.Delete)
	authGroup.GET("/me", userHandler.Me)
	authGroup.GET("/devices", userHandler.Devices)
	authGroup.DELETE("/devices/:id", userHandler.RevokeDevice)

	// ---------- GROUP ----------
	authGroup.GET("/groups", groupHandler.GetAll)
	authGroup.POST("/groups", groupHandler.Create)
	authGroup.PUT("/groups/:id", groupHandler.Update)
	authGroup.DELETE("/groups/:id", groupHandler.Delete)


	database.SyncEndpoints(db, r)
	// ================== ROOT ==================
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "server running"})
	})

	// ================== RUN ==================
	r.Run(":8080")
}
