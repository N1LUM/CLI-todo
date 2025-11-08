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

func TestTaskService_Update(t *testing.T) {
	existingTask := &models.Task{
		ID:        uuid.New(),
		Name:      "Old Name",
		CreatedAt: time.Now().Add(-48 * time.Hour),
		UpdatedAt: time.Now().Add(-48 * time.Hour),
		DeletedAt: nil,
	}

	existingTaskWithEmptyName := &models.Task{
		ID:        uuid.New(),
		Name:      "",
		CreatedAt: time.Now().Add(-48 * time.Hour),
		UpdatedAt: time.Now().Add(-48 * time.Hour),
		DeletedAt: nil,
	}

	tests := []struct {
		name      string
		task      *models.Task
		wantError bool
	}{
		{
			name:      "Valid parameters for update task",
			task:      existingTask,
			wantError: false,
		},
		{
			name:      "Invalid parameters with empty task name",
			task:      existingTaskWithEmptyName,
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, mock, service := test.CreateMockService(t)

			if !tt.wantError {
				mock.ExpectBegin()
				mock.ExpectExec(`UPDATE "tasks"`).
					WithArgs(tt.task.Name, sqlmock.AnyArg(), sqlmock.AnyArg(), tt.task.DeletedAt, tt.task.ID).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			}

			task, err := service.Task.Update(tt.task)

			if tt.wantError {
				assert.Error(t, err)
				return
			}

			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			assert.Equal(t, existingTask, task)
		})
	}
}
