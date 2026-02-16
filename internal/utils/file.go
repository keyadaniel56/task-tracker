package utils

import (
	"encoding/json"

	"os"
	"task-tracker/internal/modals"
)

func LoadTask(filename string) ([]modals.Task, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return []modals.Task{}, nil
		}
		return nil, err
	}
	var tasks []modals.Task
	err = json.Unmarshal(data, &tasks)
	return tasks, err
}
