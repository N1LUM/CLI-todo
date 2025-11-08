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

func TestTaskService_Delete(t *testing.T) {
	existingTask := &models.Task{
		ID:        uuid.New(),
		Name:      "Old Name",
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
			name:      "Valid parameters for update task",
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
				mock.ExpectBegin()
				mock.ExpectExec(`UPDATE "tasks"`).
					WithArgs(sqlmock.AnyArg(), tt.taskID).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			}

			deletedTaskID, err := service.Task.Delete(tt.taskID)

			if tt.wantError {
				assert.Error(t, err)
				return
			}

			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			assert.Equal(t, tt.taskID.String(), deletedTaskID.String())
		})
	}
}
