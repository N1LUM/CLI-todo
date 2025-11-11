package database

import (
	"CLI-todo/configs"
	"CLI-todo/internal/models"
	"fmt"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectPostgres(cfg *configs.DBConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		cfg.Host, cfg.Username, cfg.Password, cfg.DBName, cfg.Port, cfg.SSLMode,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to DB: %w", err)
	}
	return db, nil
}

func MigratePostgres(db *gorm.DB) {
	logrus.Info("Running auto-migrations...")

	if err := db.AutoMigrate(
		&models.Task{},
	); err != nil {
		logrus.Fatalf("migration failed: %v", err)
	}

	logrus.Info("Auto-migrations completed successfully")
}
