package repositories

import (
	"CLI-todo/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Task interface {
	Create(task *models.Task) error
	Update(task *models.Task) error
	GetByID(id uuid.UUID) (*models.Task, error)
	GetList() (*[]models.Task, error)
	Delete(id uuid.UUID) (uuid.UUID, error)
}

type Repository struct {
	Task *TaskRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Task: NewTaskRepository(db),
	}
}
