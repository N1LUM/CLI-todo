package services

import "CLI-todo/internal/repositories"

type TaskService struct {
	repository *repositories.TaskRepository
}

func NewTaskService(r *repositories.TaskRepository) *TaskService {
	return &TaskService{repository: r}
}
