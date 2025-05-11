package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"task-manager-app/internal/handlers"
	"task-manager-app/internal/models"
	"task-manager-app/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// setupTestServer initializes router and in-memory DB
func setupTestServer(t *testing.T) (*gin.Engine, *gorm.DB) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)
	err = db.AutoMigrate(&models.User{}, &models.Task{})
	assert.NoError(t, err)

	taskService := services.NewTaskService(db)
	taskHandler := handlers.NewTaskHandler(taskService)
	authHandlers := handlers.NewAuthHandlers(db)

	router := gin.Default()
	router.POST("/register", authHandlers.RegisterHandler)
	router.POST("/login", authHandlers.LoginHandler)

	authGroup := router.Group("/")
	authGroup.Use(handlers.AuthMiddleware())
	{
		authGroup.GET("/tasks", taskHandler.GetTasks)
		authGroup.POST("/tasks", taskHandler.CreateTask)
		authGroup.PUT("/tasks/:id", taskHandler.UpdateTask)
		authGroup.DELETE("/tasks/:id", taskHandler.DeleteTask)
	}

	return router, db
}

// helper to register user
func registerUser(t *testing.T, router *gin.Engine, username, password string) {
	body, _ := json.Marshal(gin.H{"username": username, "password": password})
	req := httptest.NewRequest("POST", "/register", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)
	assert.Equal(t, http.StatusCreated, res.Code)
}

// helper to login and get token
func loginUser(t *testing.T, router *gin.Engine, username, password string) string {
	body, _ := json.Marshal(gin.H{"username": username, "password": password})
	req := httptest.NewRequest("POST", "/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code)

	var resp map[string]interface{}
	err := json.Unmarshal(res.Body.Bytes(), &resp)
	assert.NoError(t, err)
	token, ok := resp["token"].(string)
	assert.True(t, ok)
	return token
}

func TestUserRegistrationAndLogin(t *testing.T) {
	router, _ := setupTestServer(t)
	registerUser(t, router, "testuser", "password123")

	// Duplicate register should fail
	body, _ := json.Marshal(gin.H{"username": "testuser", "password": "password123"})
	req := httptest.NewRequest("POST", "/register", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)
	assert.Equal(t, http.StatusBadRequest, res.Code)

	// Login correct credentials
	token := loginUser(t, router, "testuser", "password123")
	assert.NotEmpty(t, token)

	// Login incorrect credentials
	body, _ = json.Marshal(gin.H{"username": "testuser", "password": "wrongpass"})
	req = httptest.NewRequest("POST", "/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	res = httptest.NewRecorder()
	router.ServeHTTP(res, req)
	assert.Equal(t, http.StatusUnauthorized, res.Code)
}

func TestTaskCRUD(t *testing.T) {
	router, _ := setupTestServer(t)
	registerUser(t, router, "taskuser", "testpass")
	token := loginUser(t, router, "taskuser", "testpass")
	authHeader := "Bearer " + token

	// Create task
	taskBody, _ := json.Marshal(gin.H{
		"title":       "Test Task",
		"description": "Test Description",
		"status":      "Pending",
	})
	req := httptest.NewRequest("POST", "/tasks", bytes.NewBuffer(taskBody))
	req.Header.Set("Authorization", authHeader)
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)
	assert.Equal(t, http.StatusCreated, res.Code)
	var createdTask models.Task
	err := json.Unmarshal(res.Body.Bytes(), &createdTask)
	assert.NoError(t, err)
	assert.Equal(t, "Test Task", createdTask.Title)

	taskID := createdTask.ID

	// Get tasks
	req = httptest.NewRequest("GET", "/tasks", nil)
	req.Header.Set("Authorization", authHeader)
	res = httptest.NewRecorder()
	router.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code)
	var tasks []models.Task
	err = json.Unmarshal(res.Body.Bytes(), &tasks)
	assert.NoError(t, err)
	assert.NotEmpty(t, tasks)

	// Update task
	updateBody, _ := json.Marshal(gin.H{
		"title":       "Updated Task",
		"description": "Updated Description",
		"status":      "In-Progress",
	})
	req = httptest.NewRequest("PUT", "/tasks/"+strconv.Itoa(int(taskID)), bytes.NewBuffer(updateBody))
	req.Header.Set("Authorization", authHeader)
	req.Header.Set("Content-Type", "application/json")
	res = httptest.NewRecorder()
	router.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code)
	var updatedTask models.Task
	err = json.Unmarshal(res.Body.Bytes(), &updatedTask)
	assert.NoError(t, err)
	assert.Equal(t, "Updated Task", updatedTask.Title)

	// Delete task
	req = httptest.NewRequest("DELETE", "/tasks/"+strconv.Itoa(int(taskID)), nil)
	req.Header.Set("Authorization", authHeader)
	res = httptest.NewRecorder()
	router.ServeHTTP(res, req)
	assert.Equal(t, http.StatusNoContent, res.Code)

	// Confirm deletion
	req = httptest.NewRequest("GET", "/tasks", nil)
	req.Header.Set("Authorization", authHeader)
	res = httptest.NewRecorder()
	router.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code)
	err = json.Unmarshal(res.Body.Bytes(), &tasks)
	assert.NoError(t, err)
	assert.Len(t, tasks, 0)
}

func TestUnauthorizedAccess(t *testing.T) {
	router, _ := setupTestServer(t)
	// No auth header
	req := httptest.NewRequest("GET", "/tasks", nil)
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)
	assert.Equal(t, http.StatusUnauthorized, res.Code)

	// No auth for create
	taskBody, _ := json.Marshal(gin.H{
		"title":       "No Auth Task",
		"description": "No Auth Description",
		"status":      "Pending",
	})
	req = httptest.NewRequest("POST", "/tasks", bytes.NewBuffer(taskBody))
	req.Header.Set("Content-Type", "application/json")
	res = httptest.NewRecorder()
	router.ServeHTTP(res, req)
	assert.Equal(t, http.StatusUnauthorized, res.Code)
}
