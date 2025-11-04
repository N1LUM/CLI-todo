package services

import "CLI-todo/internal/repositories"

type Task interface{}

type Service struct {
	Task *TaskService
}

func NewService(r *repositories.Repository) *Service {
	return &Service{
		Task: NewTaskService(r.Task),
	}
}
