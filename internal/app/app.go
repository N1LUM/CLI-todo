package app

import (
	mycli "CLI-todo/cmd"
	"CLI-todo/configs"
	"CLI-todo/internal/database"
	"CLI-todo/internal/repositories"
	"CLI-todo/internal/services"
	"log"
)

func Run() {
	databaseConfig := configs.InitDatabaseConfig()

	db, err := database.ConnectPostgres(databaseConfig)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	repository := repositories.NewRepository(db)
	service := services.NewService(repository)

	mycli.Init(service)

	mycli.Execute()
}
