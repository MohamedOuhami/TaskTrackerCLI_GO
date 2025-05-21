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

// updateTaskCmd represents the updateTask command
var updateTaskCmd = &cobra.Command{
	Use:   "updateTask",
	Short: "Update the task description",
	Long:  `Update the task by giving the ID of the wanted task and the new description.`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) != 2 {
			fmt.Println("The updating command takes exactly 2 argument (1- ID of the task, 2- New task)")
			return
		}
		var tasks []models.Task
		// The arguments
		taskId := args[0]
		newDescription := args[1]

		var found = false

		// Get the list of all the tasks in the database
		tasks = utils.GetAllTasks()

		taskIdInt, err := strconv.Atoi(taskId)

		for i, task := range tasks {
			if err != nil {
				fmt.Println("Please enter a valid ID")
			}
			if task.Id == taskIdInt {
				found = true
				task.Description = newDescription

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
	rootCmd.AddCommand(updateTaskCmd)

}
