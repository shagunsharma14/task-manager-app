package services

import (
	"errors"
	"gorm.io/gorm"
	"task-manager-app/internal/models"
)

var ErrTaskNotFound = errors.New("task not found")

type TaskService struct {
	db *gorm.DB
}

func NewTaskService(db *gorm.DB) *TaskService {
	return &TaskService{db: db}
}

func (s *TaskService) GetAllTasks() ([]models.Task, error) {
	var tasks []models.Task
	if err := s.db.Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func (s *TaskService) CreateTask(task models.Task) (models.Task, error) {
	if err := s.db.Create(&task).Error; err != nil {
		return models.Task{}, err
	}
	return task, nil
}

// UpdateTask updates task by ID. Returns ErrTaskNotFound if not found.
func (s *TaskService) UpdateTask(id uint, updatedTask models.Task) (models.Task, error) {
	var task models.Task
	if err := s.db.First(&task, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Task{}, ErrTaskNotFound
		}
		return models.Task{}, err
	}
	// Update fields
	task.Title = updatedTask.Title
	task.Description = updatedTask.Description
	task.Status = updatedTask.Status
	task.DueDate = updatedTask.DueDate

	if err := s.db.Save(&task).Error; err != nil {
		return models.Task{}, err
	}
	return task, nil
}

// DeleteTask deletes task by ID. Returns ErrTaskNotFound if not found.
func (s *TaskService) DeleteTask(id uint) error {
	result := s.db.Delete(&models.Task{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ErrTaskNotFound
	}
	return nil
}
