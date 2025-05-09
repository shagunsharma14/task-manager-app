package services

import (
	"sync"
	"task-manager-app/internal/models"
)

type TaskService struct {
	mu     sync.Mutex
	tasks  []models.Task
	nextID int
}

func NewTaskService() *TaskService {
	return &TaskService{
		tasks:  []models.Task{},
		nextID: 1,
	}
}

func (s *TaskService) GetAllTasks() []models.Task {
	return s.tasks
}

func (s *TaskService) CreateTask(task models.Task) models.Task {
	task.ID = s.nextID
	s.nextID++
	s.tasks = append(s.tasks, task)
	return task
}

func (s *TaskService) UpdateTask(id int, updatedTask models.Task) (models.Task, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i, task := range s.tasks {
		if task.ID == id {
			updatedTask.ID = id
			s.tasks[i] = updatedTask
			return updatedTask, true
		}
	}
	return models.Task{}, false
}

func (s *TaskService) DeleteTask(id int) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i, task := range s.tasks {
		if task.ID == id {
			s.tasks = append(s.tasks[:i], s.tasks[i+1:]...)
			return true
		}
	}
	return false
}
