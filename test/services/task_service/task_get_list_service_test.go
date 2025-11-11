package task_service

import (
	"CLI-todo/internal/models"
	"CLI-todo/test"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestTaskService_GetList(t *testing.T) {
	existingTasks := []models.Task{
		{
			ID:        uuid.New(),
			Name:      "Task Name1",
			CreatedAt: time.Now().Add(-48 * time.Hour),
			UpdatedAt: time.Now().Add(-48 * time.Hour),
			DeletedAt: nil,
		},
		{
			ID:        uuid.New(),
			Name:      "Task Name2",
			CreatedAt: time.Now().Add(-48 * time.Hour),
			UpdatedAt: time.Now().Add(-48 * time.Hour),
			DeletedAt: nil,
		},
	}

	tests := []struct {
		name      string
		tasks     []models.Task
		wantError bool
	}{
		{
			name:      "Valid parameters for get list of tasks",
			tasks:     existingTasks,
			wantError: false,
		},
		{
			name:      "Valid parameters for get list of tasks, but tasks do not exist",
			tasks:     []models.Task{},
			wantError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, mock, service := test.CreateMockService(t)

			if !tt.wantError {
				rowBuilder := sqlmock.NewRows([]string{"id", "name", "created_at", "updated_at", "deleted_at"})
				for _, task := range tt.tasks {
					rowBuilder = rowBuilder.AddRow(task.ID, task.Name, task.CreatedAt, task.UpdatedAt, task.DeletedAt)
				}
				mock.ExpectQuery(`SELECT \* FROM \"tasks\".*`).WillReturnRows(rowBuilder)
			}

			tasks, err := service.Task.GetList()

			if tt.wantError {
				assert.Error(t, err)
				return
			}

			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			assert.Equal(t, tt.tasks, *tasks)
		})
	}
}
