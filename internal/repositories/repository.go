package repositories

import "gorm.io/gorm"

type Task interface{}

type Repository struct {
	Task *TaskRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Task: NewTaskRepository(db),
	}
}
