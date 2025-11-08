package cmd

import (
	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

var deleteTaskCmd = &cobra.Command{
	Use:   "delete <id_of_task>",
	Short: "Delete the task",
	Long:  `Delete the task by shared id`,
	Run: func(cmd *cobra.Command, args []string) {
		taskID, err := uuid.Parse(args[0])
		if err != nil {
			cmd.Printf("Failed to parse id string to uuid: %v\n", err)
			return
		}

		deletedTaskID, err := service.Task.Delete(taskID)
		if err != nil {
			cmd.Printf("Failed to delete task: %v\n", err)
			return
		}

		cmd.Printf("Task deleted! ID: %s", deletedTaskID)
	},
}
