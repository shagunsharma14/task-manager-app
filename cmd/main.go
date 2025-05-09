package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"task-manager-app/internal/handlers"
	"task-manager-app/internal/models"
	"task-manager-app/internal/services"
)

func main() {
	// Initialize Gin router
	router := gin.Default()

	// Connect to SQLite database
	db, err := gorm.Open(sqlite.Open("tasks.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Automatically perform schema migration
	if err := db.AutoMigrate(&models.Task{}); err != nil {
		log.Fatalf("Failed to migrate database schema: %v", err)
	}

	// Initialize Task Service
	taskService := services.NewTaskService(db)

	// Initialize Task Handler with the service
	taskHandler := handlers.NewTaskHandler(taskService)

	// Health check route
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "OK"})
	})

	// Task CRUD routes
	router.GET("/tasks", taskHandler.GetTasks)
	router.POST("/tasks", taskHandler.CreateTask)
	router.PUT("/tasks/:id", taskHandler.UpdateTask)
	router.DELETE("/tasks/:id", taskHandler.DeleteTask)

	// Start server on port 8080
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
