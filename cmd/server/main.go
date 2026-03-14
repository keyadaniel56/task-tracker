package main

import (
	"fmt"
	"task-tracker/internal/repository"
	"task-tracker/internal/services"
)

func main() {
	repo := repository.NewFileRepo("tasks.json")
	taskService := services.NewTaskService(repo)

	task, err := taskService.CreateTask("go shopping")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Created task: ", task)

	tasks, err := taskService.ReadAllTask()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("All tasks:", tasks)

	sortedTasks, err := taskService.SortTasksByCreatedAt()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, task := range sortedTasks {
		fmt.Println(task.ID, task.Description)
	}
}
