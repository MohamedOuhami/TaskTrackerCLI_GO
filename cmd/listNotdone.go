/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/MohamedOuhami/TaskTrackerCLI_GO/models"
	"github.com/MohamedOuhami/TaskTrackerCLI_GO/utils"
	"github.com/spf13/cobra"
)

// listNotdoneCmd represents the listNotdone command
var listNotdoneCmd = &cobra.Command{
	Use:   "listNotdone",
	Short: "A command to list all of the Not Done tasks",
	Long:  `A command to list all of the Not Done tasks.`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) != 0 {
			fmt.Println("The listing command does not accept any argument")
			return
		}

		utils.CheckIfFilExists()

		var tasks []models.Task = utils.GetAllTasks()

		var InProgressTasks []models.Task = []models.Task{}

		for _, task := range tasks {
			if task.Status == "Not done" {
				InProgressTasks = append(InProgressTasks, task)
			}
		}

		utils.PrintDataInTable(InProgressTasks)

	},
}

func init() {
	rootCmd.AddCommand(listNotdoneCmd)
}
