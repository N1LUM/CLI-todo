package test

import (
	"CLI-todo/internal/repositories"
	"CLI-todo/internal/services"
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func CreateMockService(t *testing.T) (*sql.DB, sqlmock.Sqlmock, *services.Service) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create sqlmock: %v", err)
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to open gorm db: %v", err)
	}

	repository := repositories.NewRepository(gormDB)
	service := services.NewService(repository)

	return db, mock, service
}
