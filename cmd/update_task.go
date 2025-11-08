package cmd

import (
	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

var updateTaskCmd = &cobra.Command{
	Use:   "update <id_of_task> <name_of_task>",
	Short: "Update the task",
	Long:  `Update the task with shared id`,
	Run: func(cmd *cobra.Command, args []string) {
		taskID, err := uuid.Parse(args[0])
		if err != nil {
			cmd.Printf("Failed to parse id string to uuid: %v\n", err)
			return
		}
		taskName := args[1]

		task, err := service.Task.GetByID(taskID)
		if err != nil {
			cmd.Printf("Failed to get task by id: %v\n", err)
			return
		}

		task.Name = taskName

		updatedTask, err := service.Task.Update(task)
		if err != nil {
			cmd.Printf("Failed to update task: %v\n", err)
			return
		}

		cmd.Printf("Task updated! ID: %s, Name: %s\n", updatedTask.ID, updatedTask.Name)
	},
}
