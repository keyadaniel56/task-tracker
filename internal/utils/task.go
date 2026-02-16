package utils

import (
	"encoding/json"
	"os"
	"task-tracker/internal/modals"
)

func SaveTasks(filename string, tasks []modals.Task) error {
	data, err := json.MarshalIndent(tasks, "", "")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}
