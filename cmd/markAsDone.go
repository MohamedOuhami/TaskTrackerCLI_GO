/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"strconv"

	"github.com/MohamedOuhami/TaskTrackerCLI_GO/models"
	"github.com/MohamedOuhami/TaskTrackerCLI_GO/utils"
	"github.com/spf13/cobra"
)

// markAsDoneCmd represents the markAsDone command
var markAsDoneCmd = &cobra.Command{
	Use:   "markAsDone",
	Short: "A command to mark a task as being Done",
	Long:  `A command to mark a task as being done.`,
	Run: func(cmd *cobra.Command, args []string) {

		var tasks []models.Task
		// The arguments
		taskId := args[0]

		var found = false

		// Get the list of all the tasks in the database
		tasks = utils.GetAllTasks()

		taskIdInt, err := strconv.Atoi(taskId)

		for i, task := range tasks {
			if err != nil {
				panic("error converting ID to int")
			}
			if task.Id == taskIdInt {
				found = true
				task.Status = "Done"

				tasks[i] = task

				utils.SaveNewTasks(tasks)

			}
		}

		if found {
			println("Found the task and updated It")
		} else {
			println("The task of id", taskIdInt, "is not found")
		}
	},
}

func init() {
	rootCmd.AddCommand(markAsDoneCmd)
}
