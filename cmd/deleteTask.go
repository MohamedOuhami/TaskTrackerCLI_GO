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

// deleteTaskCmd represents the deleteTask command
var deleteTaskCmd = &cobra.Command{
	Use:   "deleteTask",
	Short: "Deleting the Task by providing the ID of the task",
	Long:  `A function to delete a task by giving Its ID and remove It from the database.`,
	Run: func(cmd *cobra.Command, args []string) {

		var tasks []models.Task

		var found = false

		// Check if the file exists
		utils.CheckIfFilExists()

		tasks = utils.GetAllTasks()
		// Getting the id of the task to delete
		if idInt, convErr := strconv.Atoi(args[0]); convErr != nil {
			panic(convErr)

		} else {

			// Get the ATM by ID
			for i, task := range tasks {

				if task.Id == idInt {

					found = true
					fmt.Println("Found the task")
					var newTaskList = append(tasks[:i], tasks[i+1:]...)
					utils.SaveNewTasks(newTaskList)
				}
			}

			if found {
				fmt.Println("Deleted the task of id ", idInt)
			} else {
				fmt.Println("The task of id", idInt, "does not exist")
			}

		}
	},
}

func init() {
	rootCmd.AddCommand(deleteTaskCmd)

}
