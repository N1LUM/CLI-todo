package app

import (
	mycli "CLI-todo/cmd"
	"CLI-todo/internal/repositories"
	"CLI-todo/internal/services"
)

func Run() {
	repository := repositories.NewRepository()
	service := services.NewService(repository)

	mycli.Init(service)

	mycli.Execute()
}
