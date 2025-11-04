package cmd

import "github.com/spf13/cobra"

var addTaskCmd = &cobra.Command{
	Use:   "add <name_of_task>",
	Short: "Create a new task",
	Long:  `Create a new task with shared name`,
	Run: func(cmd *cobra.Command, args []string) {
		//taskName := args[0]
		//service.Task.Create(taskName)
	},
}
