package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"text/tabwriter"

	"github.com/MohamedOuhami/TaskTrackerCLI_GO/models"
)

func CheckFileError(e error) {
	if e != nil {
		panic(e)
	}
}

func marshalTasks(tasks []models.Task) string {

	newTaskJSON, _ := json.MarshalIndent(tasks, "", " ")
	newTaskContent := string(newTaskJSON)

	return newTaskContent

}

func SaveNewTasks(newTasks []models.Task) {

	f, err := os.OpenFile("./tasklist.json", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}

	defer f.Close()
	_, writeErr := f.Write([]byte(marshalTasks(newTasks)))

	if writeErr != nil {
		panic(err)
	}

}

func PrintDataInTable(tasks []models.Task) {

	// Printing the data in tabular manner
	data := [][]string{
		{"ID", "Description", "Status"},
	}

	for _, task := range tasks {
		data = append(data, []string{strconv.Itoa(task.Id), task.Description, task.Status})
	}

	fmt.Println("\n")
	fmt.Println("Tasks count : ", len(tasks))
	fmt.Println("==============================")
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	for _, row := range data {
		fmt.Fprintf(w, "%s\t%s\t%s\t\n", row[0], row[1], row[2])
	}
	w.Flush()
}
func CheckIfFilExists() {
	// Create the tasklist.json file if It does not exist
	if _, err := os.Stat("./tasklist.json"); err != nil {

		os.Create("./tasklist.json")
	}

}

// Get all the Tasks in the JSON
func GetAllTasks() []models.Task {

	var tasks []models.Task
	// Open the json file
	if data, readErr := os.ReadFile("./tasklist.json"); readErr != nil {
		panic(readErr)
	} else {
		tasksData := string(data)
		if unmarshalError := json.Unmarshal([]byte(tasksData), &tasks); unmarshalError != nil {
			panic(unmarshalError)
		}
		return tasks
	}

}
