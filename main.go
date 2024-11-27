package main

import (
	"fmt"
)

func main() {
	// Create a slice to store the tasks
	var tasks []string

	// Welcome message
	fmt.Println("Welcome to the To-Do List Application!")

	// Main loop to manage tasks
	for {
		// Display the menu
		fmt.Println("\nPlease choose an option:")
		fmt.Println("1. Add a task")
		fmt.Println("2. View tasks")
		fmt.Println("3. Remove a task")
		fmt.Println("4. Exit")

		// Get user input
		var choice int
		fmt.Scan(&choice)

		// Perform actions based on the choice
		switch choice {
		case 1:
			// Add a task
			var task string
			fmt.Print("Enter the task: ")
			fmt.Scan(&task)
			tasks = append(tasks, task)
			fmt.Println("Task added!")

		case 2:
			// View tasks
			if len(tasks) == 0 {
				fmt.Println("No tasks in the list.")
			} else {
				fmt.Println("\nYour tasks:")
				for i, task := range tasks {
					fmt.Printf("%d. %s\n", i+1, task)
				}
			}

		case 3:
			// Remove a task
			if len(tasks) == 0 {
				fmt.Println("No tasks to remove.")
			} else {
				var taskNumber int
				fmt.Println("Enter the task number to remove:")
				fmt.Scan(&taskNumber)
				// Check if the input is valid
				if taskNumber < 1 || taskNumber > len(tasks) {
					fmt.Println("Invalid task number.")
				} else {
					// Remove the task
					tasks = append(tasks[:taskNumber-1], tasks[taskNumber:]...)
					fmt.Println("Task removed!")
				}
			}

		case 4:
			// Exit the application
			fmt.Println("Exiting the To-Do List Application. Goodbye!")
			return

		default:
			// Handle invalid input
			fmt.Println("Invalid choice. Please choose a valid option.")
		}
	}
}
