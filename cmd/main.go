package main

import (
	"log"
	"task-manager-app/internal/handlers"
	"task-manager-app/internal/models"
	"task-manager-app/internal/services"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	router := gin.Default()

	// Configure CORS to allow frontend requests
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // Frontend origin
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	db, err := gorm.Open(sqlite.Open("tasks.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect database: %v", err)
	}

	// Migrate schema
	if err := db.AutoMigrate(&models.User{}, &models.Task{}); err != nil {
		log.Fatalf("Failed to migrate database schema: %v", err)
	}

	taskService := services.NewTaskService(db)
	taskHandler := handlers.NewTaskHandler(taskService)

	authHandlers := handlers.NewAuthHandlers(db)

	// Public routes
	router.POST("/register", authHandlers.RegisterHandler)
	router.POST("/login", authHandlers.LoginHandler)

	// Protected task routes
	protected := router.Group("/")
	protected.Use(handlers.AuthMiddleware())
	{
		protected.GET("/tasks", taskHandler.GetTasks)
		protected.POST("/tasks", taskHandler.CreateTask)
		protected.PUT("/tasks/:id", taskHandler.UpdateTask)
		protected.DELETE("/tasks/:id", taskHandler.DeleteTask)
	}

	// Health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "OK"})
	})

	log.Println("Starting server on :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
