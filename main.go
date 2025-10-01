package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"
)

const (
	StatusTodo       = "todo"
	StatusInProgress = "in-progress"
	StatusDone       = "done"
)

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

const fileName = "tasks.json"

func loadTasks() ([]Task, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			return []Task{}, nil
		}
		return nil, err
	}

	var tasks []Task
	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}

func saveTasks(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, data, 0644)
}

func addTask(description string) {
	tasks, err := loadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	newID := 1
	if len(tasks) > 0 {
		newID = tasks[len(tasks)-1].ID + 1
	}

	now := time.Now()
	newTask := Task{
		ID:          newID,
		Description: description,
		Status:      StatusTodo,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	tasks = append(tasks, newTask)
	if err := saveTasks(tasks); err != nil {
		fmt.Println("Error saving task:", err)
		return
	}

	fmt.Printf("Task '%s' added successfully with ID: %d\n", newTask.Description, newTask.ID)
}

func listTasks() {
	tasks, err := loadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks found. Use the 'add' command to create one.")
		return
	}

	fmt.Println("--- Task List ---")
	for _, t := range tasks {
		fmt.Printf("ID: %d\n", t.ID)
		fmt.Printf("  Description: %s\n", t.Description)
		fmt.Printf("  Status: %s\n", t.Status)
		fmt.Printf("  Created At: %s\n", t.CreatedAt.Format("2006-01-02 15:04:05"))
		fmt.Printf("  Updated At: %s\n\n", t.UpdatedAt.Format("2006-01-02 15:04:05"))
	}
}

func updateStatus(id int, newStatus string) {
	if newStatus != StatusTodo && newStatus != StatusInProgress && newStatus != StatusDone {
		fmt.Printf("Invalid status. Please use one of: %s, %s, %s\n", StatusTodo, StatusInProgress, StatusDone)
		return
	}

	tasks, err := loadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	found := false
	for i, t := range tasks {
		if t.ID == id {
			tasks[i].Status = newStatus
			tasks[i].UpdatedAt = time.Now()
			found = true
			break
		}
	}

	if !found {
		fmt.Println("Task with the provided ID not found.")
		return
	}

	if err := saveTasks(tasks); err != nil {
		fmt.Println("Error saving tasks:", err)
		return
	}
	fmt.Printf("Status of task %d updated to '%s'!\n", id, newStatus)
}

func deleteTask(id int) {
	tasks, err := loadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	indexToRemove := -1
	for i, t := range tasks {
		if t.ID == id {
			indexToRemove = i
			break
		}
	}

	if indexToRemove == -1 {
		fmt.Println("Task with the provided ID not found.")
		return
	}

	tasks = append(tasks[:indexToRemove], tasks[indexToRemove+1:]...)

	if err := saveTasks(tasks); err != nil {
		fmt.Println("Error saving tasks:", err)
		return
	}
	fmt.Println("Task deleted successfully!")
}

func showHelp() {
	fmt.Println("Usage: tasks <command> [arguments]")
	fmt.Println("\nAvailable commands:")
	fmt.Println("  list                           - Displays all tasks.")
	fmt.Println("  add \"<description>\"            - Adds a new task.")
	fmt.Println("  update <ID> <new-status>     - Updates a task's status (todo, in-progress, done).")
	fmt.Println("  delete <ID>                    - Deletes a task.")
}

func main() {
	if len(os.Args) < 2 {
		showHelp()
		return
	}

	command := os.Args[1]

	switch command {
	case "list":
		listTasks()
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Error: Please provide a description for the task.")
			fmt.Println("Usage: tasks add \"<task description>\"")
			return
		}
		addTask(os.Args[2])
	case "update":
		if len(os.Args) < 4 {
			fmt.Println("Error: Please provide the task ID and the new status.")
			fmt.Println("Usage: tasks update <ID> <new-status>")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("The ID must be a number.")
			return
		}
		newStatus := os.Args[3]
		updateStatus(id, newStatus)
	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Error: Please provide the ID of the task to delete.")
			fmt.Println("Usage: tasks delete <ID>")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("The ID must be a number.")
			return
		}
		deleteTask(id)
	default:
		fmt.Println("Unknown command:", command)
		showHelp()
	}
}
