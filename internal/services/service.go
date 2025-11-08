package services

import (
	"CLI-todo/internal/models"
	"CLI-todo/internal/repositories"
)

type Task interface {
	Create(name string) (*models.Task, error)
	Update(name string) (*models.Task, error)
}

type Service struct {
	Task *TaskService
}

func NewService(r *repositories.Repository) *Service {
	return &Service{
		Task: NewTaskService(r.Task),
	}
}
