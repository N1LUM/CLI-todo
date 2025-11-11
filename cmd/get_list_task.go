package cmd

import (
	"reflect"

	"github.com/spf13/cobra"
)

var getListTaskCmd = &cobra.Command{
	Use:   "get-list",
	Short: "Get list of tasks",
	Long:  `Get list of tasks in descending order by created_at`,
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := service.Task.GetList()
		if err != nil {
			cmd.Printf("Failed to get all tasks: %v\n", err)
			return
		}

		cmd.Print("Got list of tasks: \n")

		for index, task := range *tasks {
			cmd.Printf("\nTask #%d\n", index+1)

			taskValue := reflect.ValueOf(task)
			taskType := reflect.TypeOf(task)

			for i := 0; i < taskType.NumField(); i++ {
				fieldInfo := taskType.Field(i)
				fieldValue := taskValue.Field(i)

				cmd.Printf("%s: %v\n", fieldInfo.Name, fieldValue.Interface())
			}
		}
	},
}
