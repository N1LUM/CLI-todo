package services

import (
	"CLI-todo/internal/models"
	"CLI-todo/internal/repositories"
)

type TaskService struct {
	repository *repositories.TaskRepository
}

func NewTaskService(r *repositories.TaskRepository) *TaskService {
	return &TaskService{repository: r}
}

func (service *TaskService) Create(name string) (*models.Task, error) {
	task := models.Task{Name: name}

	if err := service.repository.Create(&task); err != nil {
		return nil, err
	}

	return &task, nil
}
