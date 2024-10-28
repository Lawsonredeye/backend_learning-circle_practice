/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type Task struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Done bool `json:"done"`
}

// LoadTask retrive all added / stored tasks from the file storage.
func loadTasks() ([]Task, error) {
	file, err := os.ReadFile("tasks.json")
	if err != nil {
		if os.IsNotExist(err) {
			return []Task{}, nil
		}
		return nil, err
	}

	if len(file) == 0 {
		return []Task{}, nil
	}

	var tasks []Task
	if err := json.Unmarshal(file, &tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}

// SaveTasks saves stores the newly added tasks to the file storage
func saveTasks(tasks []Task) error {
	data, err := json.Marshal(tasks)
	if err != nil {
		return err
	}
	return os.WriteFile("tasks.json", data, 0644)
}

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [task name]",
	Short: "Add a new task to the local storage with a new ID",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		title := args[0]
		tasks, err := loadTasks()
		if err != nil {
			fmt.Println("Error loading tasks:", err)
			return
		}
		
		var id int
		if len(tasks) == 0 {
			id = 1
		} else {
			id = tasks[len(tasks) - 1].ID + 1
		}

		newTask := Task {
			ID: id,
			Title: title,
			Done: false,
		}
		tasks = append(tasks, newTask)

		if err := saveTasks(tasks); err != nil {
			fmt.Println("Error saving tasks:", err)
			return
		}

		fmt.Printf("Added task: %s\n", title)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
