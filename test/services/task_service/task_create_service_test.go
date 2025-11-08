package task_service

import (
	"CLI-todo/test"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestTaskService_Create(t *testing.T) {
	tests := []struct {
		name      string
		taskName  string
		wantError bool
	}{
		{
			name:      "Valid parameters for create task",
			taskName:  "New Task",
			wantError: false,
		},
		{
			name:      "Invalid parameters with empty task name",
			taskName:  "",
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, mock, service := test.CreateMockService(t)

			if !tt.wantError {
				mock.ExpectBegin()
				mock.ExpectQuery(`INSERT INTO "tasks"`).
					WithArgs(tt.taskName, sqlmock.AnyArg(), sqlmock.AnyArg(), nil).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(uuid.New()))
				mock.ExpectCommit()
			}

			task, err := service.Task.Create(tt.taskName)

			if tt.wantError {
				assert.Error(t, err)
				return
			}

			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			assert.IsType(t, uuid.UUID{}, task.ID)
			assert.WithinDurationf(t, time.Now(), task.CreatedAt, time.Second, "CreatedAt должна быть близка к %v", time.Now())
			assert.WithinDurationf(t, time.Now(), task.UpdatedAt, time.Second, "UpdatedAt должна быть близка к %v", time.Now())
		})
	}
}
