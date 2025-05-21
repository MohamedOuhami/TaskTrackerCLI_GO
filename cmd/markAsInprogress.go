/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/MohamedOuhami/TaskTrackerCLI_GO/models"
	"github.com/MohamedOuhami/TaskTrackerCLI_GO/utils"
	"github.com/spf13/cobra"
)

// markAsInprogressCmd represents the markAsInprogress command
var markAsInprogressCmd = &cobra.Command{
	Use:   "markAsInprogress",
	Short: "A command to mark a task as In progress",
	Long:  `A command to mark a task as In progress.`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) != 1 {
			fmt.Println("The marking command only takes one argument (ID of the task)")
		}
		var tasks []models.Task
		// The arguments
		taskId := args[0]

		var found = false

		// Get the list of all the tasks in the database
		tasks = utils.GetAllTasks()

		taskIdInt, err := strconv.Atoi(taskId)

		for i, task := range tasks {
			if err != nil {
				fmt.Println("Please enter a valid ID")
				return
			}
			if task.Id == taskIdInt {
				found = true
				task.Status = "In Progress"

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
	rootCmd.AddCommand(markAsInprogressCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// markAsInprogressCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// markAsInprogressCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
