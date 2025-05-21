/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/MohamedOuhami/TaskTrackerCLI_GO/models"
	"github.com/MohamedOuhami/TaskTrackerCLI_GO/utils"
	"github.com/spf13/cobra"
)

// listDoneCmd represents the listDone command
var listDoneCmd = &cobra.Command{
	Use:   "listDone",
	Short: "A command to list the Done tasks",
	Long:  `A command to list the Done tasks.`,
	Run: func(cmd *cobra.Command, args []string) {

		utils.CheckIfFilExists()

		var tasks []models.Task = utils.GetAllTasks()

		var InProgressTasks []models.Task = []models.Task{}

		for _, task := range tasks {
			if task.Status == "Done" {
				InProgressTasks = append(InProgressTasks, task)
			}
		}

		utils.PrintDataInTable(InProgressTasks)

	},
}

func init() {
	rootCmd.AddCommand(listDoneCmd)

}
