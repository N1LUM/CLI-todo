package repositories

import (
	"CLI-todo/internal/models"

	"gorm.io/gorm"
)

type Task interface {
	Create(task *models.Task) error
}

type Repository struct {
	Task *TaskRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Task: NewTaskRepository(db),
	}
}
