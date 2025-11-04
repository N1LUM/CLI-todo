package cmd

import (
	"CLI-todo/internal/services"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "CLI-todo",
	Short: "A brief description of your application",
	Long:  `A longer description that spans multiple lines and likely contains examples and usage of using your application`,
}

var service *services.Service

func Init(s *services.Service) {
	service = s
}

// Execute запускает rootCmd
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

// init() — сюда добавляем только флаги, глобальные настройки и подкоманды
func init() {
	rootCmd.AddCommand(addTaskCmd)
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
