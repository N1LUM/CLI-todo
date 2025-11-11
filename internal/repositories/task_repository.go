package repositories

import (
	"CLI-todo/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TaskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{
		db: db,
	}
}

func (r *TaskRepository) Create(task *models.Task) error {
	if err := r.db.Create(task).Error; err != nil {
		return err
	}
	return nil
}

func (r *TaskRepository) Update(task *models.Task) error {
	if err := r.db.Save(task).Error; err != nil {
		return err
	}
	return nil
}

func (r *TaskRepository) GetByID(id uuid.UUID) (*models.Task, error) {
	var task models.Task
	if err := r.db.First(&task, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &task, nil
}

func (r *TaskRepository) GetList() (*[]models.Task, error) {
	var tasks []models.Task
	if err := r.db.Order("created_at desc").Find(&tasks).Error; err != nil {
		return nil, err
	}
	return &tasks, nil
}

func (r *TaskRepository) Delete(id uuid.UUID) error {
	if err := r.db.Delete(&models.Task{}, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}
