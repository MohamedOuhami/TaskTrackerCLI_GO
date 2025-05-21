/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/MohamedOuhami/TaskTrackerCLI_GO/models"
	"github.com/MohamedOuhami/TaskTrackerCLI_GO/utils"
	"github.com/spf13/cobra"
)

// listAllCmd represents the listAll command
var listAllCmd = &cobra.Command{
	Use:   "listAll",
	Short: "A command to list all the tasks in the database",
	Long:  `A command to list all of the tasks in the database. All statuses are included.`,
	Run: func(cmd *cobra.Command, args []string) {

		// Checking if the file exists
		utils.CheckIfFilExists()

		var tasks []models.Task = utils.GetAllTasks()
		// Print the data in the CLI
		utils.PrintDataInTable(tasks)
	},
}

func init() {
	rootCmd.AddCommand(listAllCmd)
}
