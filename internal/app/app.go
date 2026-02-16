package app

import (
	"log"
	"os"
	"time"

	"task-tracker/internal/modals"
	"task-tracker/internal/utils"

	"github.com/google/uuid"
)

func App() {
	if len(os.Args) < 2 {
		utils.Type("Usage: task [add|update|delete]", 3*time.Millisecond)
		return
	}

	command := os.Args[1]
	filename := "tasks.json"

	tasks, err := utils.LoadTask(filename)
	if err != nil {
		log.Fatal(err)
	}

	switch command {

	case "add":
		if len(os.Args) < 3 {
			utils.Type("Usage: task add \"description\"", 3*time.Millisecond)
			return
		}

		description := os.Args[2]

		newTask := modals.Task{
			ID:          uuid.New().String(),
			Description: description,
			Completed:   false,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		tasks = append(tasks, newTask)

		err = utils.SaveTasks(filename, tasks)
		if err != nil {
			log.Fatal(err)
		}

		utils.Type("Task added successfully!", 3*time.Millisecond)

	default:
		utils.Type("Unknown command", 3*time.Millisecond)
	}
}
