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

func (s *TaskService) GetTasksByUsername(username string) ([]models.Task, error) {
	var user models.User
	if err := s.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}

	var tasks []models.Task
	if err := s.db.Where("user_id = ?", user.ID).Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func (s *TaskService) CreateTaskForUser(username string, task models.Task) (models.Task, error) {
	var user models.User
	if err := s.db.Where("username = ?", username).First(&user).Error; err != nil {
		return models.Task{}, err
	}
	task.UserID = user.ID
	if err := s.db.Create(&task).Error; err != nil {
		return models.Task{}, err
	}
	return task, nil
}

func (s *TaskService) UpdateTaskForUser(username string, id uint, updatedTask models.Task) (models.Task, error) {
	var user models.User
	if err := s.db.Where("username = ?", username).First(&user).Error; err != nil {
		return models.Task{}, err
	}

	var task models.Task
	if err := s.db.First(&task, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Task{}, ErrTaskNotFound
		}
		return models.Task{}, err
	}

	if task.UserID != user.ID {
		return models.Task{}, errors.New("not authorized to update this task")
	}

	task.Title = updatedTask.Title
	task.Description = updatedTask.Description
	task.Status = updatedTask.Status
	task.DueDate = updatedTask.DueDate

	if err := s.db.Save(&task).Error; err != nil {
		return models.Task{}, err
	}
	return task, nil
}

func (s *TaskService) DeleteTaskForUser(username string, id uint) error {
	var user models.User
	if err := s.db.Where("username = ?", username).First(&user).Error; err != nil {
		return err
	}

	var task models.Task
	if err := s.db.First(&task, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrTaskNotFound
		}
		return err
	}

	if task.UserID != user.ID {
		return errors.New("not authorized to delete this task")
	}

	result := s.db.Delete(&task)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ErrTaskNotFound
	}
	return nil
}
