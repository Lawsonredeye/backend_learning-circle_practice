package cmd

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

func deleteTask(taskId int) error {
	// search for file first of all
	// if file is not located log the error to the console
	// search for the id of the task to be deleted, if its not found,
	// then return an error
	// if its found delete the found task from the file storage
	tasks, err := loadTasks()

	if err != nil {
		return errors.New("File not found")
	}

	found := false

	var tempSlice []Task
	for _, task := range tasks {
		if task.ID == taskId {
			found = true
		} else {
			tempSlice = append(tempSlice, task)
		}
	}

	if err = saveTasks(tempSlice); err != nil {
		return errors.New("error occured while saving file")
	}
	if found {
		fmt.Println("Task found and deleted successfully")
	}
	return nil
}

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete [task id]",
	Short: "deletes tasks from the file storage list",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		taskId, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("error converting to string")
			return
		}

		if err = deleteTask(taskId); err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
