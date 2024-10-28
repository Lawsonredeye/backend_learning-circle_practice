package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "go run main.go list",
	Short: "list all saved task on the local storage to the console",
	Run: func(cmd *cobra.Command, args []string) {
		list()	
	},
}

func list() {
	// lists all the json value on the file storage.
	// open the file, then convert the data from JSON into a go object
	// then print to the console the fetched data.
	file, err := os.ReadFile("tasks.json")
	if err != nil {
		if os.IsNotExist(err) {
			return
		}
		return
	}
	
	// check if file is empty
	if len(file) == 0 {
		return
	}

	var allTasks []Task
	if err = json.Unmarshal(file, &allTasks); err != nil {
		return
	}

	fmt.Printf("%+v\n", allTasks)
}

func init() {
	rootCmd.AddCommand(listCmd)
}
