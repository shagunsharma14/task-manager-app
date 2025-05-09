package main

import (
	"github.com/gin-gonic/gin"
	"task-manager-app/internal/handlers"
	"task-manager-app/internal/services"
)

func main() {
	router := gin.Default()

	// Initialize Task Service
	taskService := services.NewTaskService()
	// Initialize Task Handler with service
	taskHandler := handlers.NewTaskHandler(taskService)

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "OK"})
	})

	// Task endpoints
	router.GET("/tasks", taskHandler.GetTasks)
	router.POST("/tasks", taskHandler.CreateTask)
	router.PUT("/tasks/:id", taskHandler.UpdateTask)
	router.DELETE("/tasks/:id", taskHandler.DeleteTask)

	// Start server on port 8080
	router.Run(":8080")
}
