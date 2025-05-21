/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"time"

	"github.com/MohamedOuhami/TaskTrackerCLI_GO/models"
	"github.com/MohamedOuhami/TaskTrackerCLI_GO/utils"
	"github.com/spf13/cobra"
)

// addTaskCmd represents the addTask command
var addTaskCmd = &cobra.Command{
	Use:   "addTask",
	Short: "Add a new task to the JSON file",
	Long:  `Create a new task and add it to the tasks JSON file`,
	Run: func(cmd *cobra.Command, args []string) {

		// Getting the task name from the arguments
		taskName := args[0]
		// Get the existing data in the json
		var tasks []models.Task

		tasks = utils.GetAllTasks()

		var lastId int

		// Now, I need a way to get the latest id
		if len(tasks) > 0 {

			lastId = tasks[len(tasks)-1].Id
		} else {
			lastId = 0
		}

		// Here, we need to create a Task out of the content we got
		newTask := models.Task{Id: lastId + 1, Description: taskName, Status: "Not done", CreatedAt: time.Now(), UpdatedAt: time.Now()}

		tasks = append(tasks, newTask)

		// Save the new content
		utils.SaveNewTasks(tasks)
	},
}

func init() {
	rootCmd.AddCommand(addTaskCmd)
}
