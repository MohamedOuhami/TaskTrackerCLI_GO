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

// markAsNotDoneCmd represents the markAsNotDone command
var markAsNotDoneCmd = &cobra.Command{
	Use:   "markAsNotDone",
	Short: "A command to mark a task as not done",
	Long:  `A command to mark a task as not done.`,
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
				task.Status = "Not done"

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
	rootCmd.AddCommand(markAsNotDoneCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// markAsNotDoneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// markAsNotDoneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
