package cmd

import "github.com/spf13/cobra"

var addTaskCmd = &cobra.Command{
	Use:   "add <name_of_task>",
	Short: "Create a new task",
	Long:  `Create a new task with shared name`,
	Run: func(cmd *cobra.Command, args []string) {
		taskName := args[0]
		task, err := service.Task.Create(taskName)
		if err != nil {
			cmd.Printf("Failed to create task: %v\n", err)
			return
		}

		cmd.Printf("Task created! ID: %s, Name: %s\n", task.ID, task.Name)
	},
}
