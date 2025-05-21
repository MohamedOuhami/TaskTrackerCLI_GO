/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/MohamedOuhami/TaskTrackerCLI_GO/models"
	"github.com/MohamedOuhami/TaskTrackerCLI_GO/utils"
	"github.com/spf13/cobra"
)

// listInprogressCmd represents the listInprogress command
var listInprogressCmd = &cobra.Command{
	Use:   "listInprogress",
	Short: "A command to print the task that are in progress",
	Long:  `A command to print the tasks that are still In progress`,
	Run: func(cmd *cobra.Command, args []string) {

		utils.CheckIfFilExists()

		var tasks []models.Task = utils.GetAllTasks()

		var InProgressTasks []models.Task = []models.Task{}

		for _, task := range tasks {
			if task.Status == "In Progress" {
				InProgressTasks = append(InProgressTasks, task)
			}
		}

		utils.PrintDataInTable(InProgressTasks)

	},
}

func init() {
	rootCmd.AddCommand(listInprogressCmd)
}
