package services

import (
	"CLI-todo/internal/models"
	"CLI-todo/internal/repositories"
	"errors"
)

type TaskService struct {
	repository *repositories.TaskRepository
}

func NewTaskService(r *repositories.TaskRepository) *TaskService {
	return &TaskService{repository: r}
}

func (service *TaskService) Create(name string) (*models.Task, error) {
	if len(name) == 0 {
		return nil, errors.New("Name must be longer than 0 symbols")
	}

	task := models.Task{Name: name}

	if err := service.repository.Create(&task); err != nil {
		return nil, err
	}

	return &task, nil
}
