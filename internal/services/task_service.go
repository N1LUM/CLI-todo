package services

import (
	"CLI-todo/internal/models"
	"CLI-todo/internal/repositories"
	"errors"

	"github.com/google/uuid"
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

func (service *TaskService) Update(task *models.Task) (*models.Task, error) {
	if len(task.Name) == 0 {
		return nil, errors.New("New name must be longer than 0 symbols")
	}

	if err := service.repository.Update(task); err != nil {
		return nil, err
	}

	return task, nil
}

func (service *TaskService) GetByID(id uuid.UUID) (*models.Task, error) {
	return service.repository.GetByID(id)
}

func (service *TaskService) GetList() (*[]models.Task, error) {
	return service.repository.GetList()
}

func (service *TaskService) Delete(id uuid.UUID) (uuid.UUID, error) {
	if err := service.repository.Delete(id); err != nil {
		return uuid.Nil, err
	}
	return id, nil
}
