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

func TestTaskService_GetByID(t *testing.T) {
	existingTask := &models.Task{
		ID:        uuid.New(),
		Name:      "Task Name",
		CreatedAt: time.Now().Add(-48 * time.Hour),
		UpdatedAt: time.Now().Add(-48 * time.Hour),
		DeletedAt: nil,
	}

	tests := []struct {
		name      string
		taskID    uuid.UUID
		wantError bool
	}{
		{
			name:      "Valid parameters for get task by id",
			taskID:    existingTask.ID,
			wantError: false,
		},
		{
			name:      "Valid parameters when task doesn't exist",
			taskID:    uuid.New(),
			wantError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, mock, service := test.CreateMockService(t)

			if !tt.wantError {
				mock.ExpectQuery(`SELECT \* FROM \"tasks\".*`).
					WithArgs(tt.taskID, sqlmock.AnyArg()).
					WillReturnRows(sqlmock.NewRows([]string{"id", "name", "created_at", "updated_at", "deleted_at"}).
						AddRow(tt.taskID, existingTask.Name, existingTask.CreatedAt, existingTask.UpdatedAt, existingTask.DeletedAt))
			}

			task, err := service.Task.GetByID(tt.taskID)

			if tt.wantError {
				assert.Error(t, err)
				return
			}

			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			assert.Equal(t, tt.taskID.String(), task.ID.String())
		})
	}
}
