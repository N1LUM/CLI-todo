package cmd

import (
	"reflect"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

var getTaskByIdCmd = &cobra.Command{
	Use:   "get-task-by-id <task_id>",
	Short: "Get task by id",
	Long:  `Get the task by shared id`,
	Run: func(cmd *cobra.Command, args []string) {
		taskID, err := uuid.Parse(args[0])
		if err != nil {
			cmd.Printf("Failed to parse id string to uuid: %v\n", err)
			return
		}

		task, err := service.Task.GetByID(taskID)
		if err != nil {
			cmd.Printf("Failed to get task by id: %v\n", err)
			return
		}

		cmd.Printf("Got task by id: %v\n\n", taskID)

		taskValue := reflect.ValueOf(task)
		taskType := reflect.TypeOf(task)

		if taskType.Kind() == reflect.Ptr {
			taskType = taskType.Elem()
			taskValue = taskValue.Elem()
		}

		for i := 0; i < taskType.NumField(); i++ {
			fieldInfo := taskType.Field(i)
			fieldValue := taskValue.Field(i)

			cmd.Printf("%s: %v\n", fieldInfo.Name, fieldValue.Interface())
		}
	},
}
