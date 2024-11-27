# To-Do List Application 📝

A simple command-line To-Do List Application written in Go. This app allows users to manage tasks by adding, viewing, and removing tasks. It's a beginner-friendly project to help learn basic Go programming concepts.

You can find the project repository here: [To-Do List Application - GitHub](https://github.com/n1tik91/To-Do-List-)

---

## 📖 Overview

The **To-Do List Application** is a basic task manager built in Go that lets you:
- Add tasks to a list.
- View all the tasks in the list.
- Remove tasks by selecting a task number.
- Exit the application when done.

The project uses fundamental Go concepts like loops, conditionals, slices, and user input handling.

---

## 🚀 Features

- Add new tasks to the to-do list.
- View all tasks with a numbered list.
- Remove tasks by their number.
- Simple text-based menu for easy navigation.
- Exit the application when finished.

---

## 🛠️ Requirements

- **Go**: You need Go installed on your system. You can download it from [Go Downloads](https://golang.org/dl/).

---

## 💻 How to Run

1. **Clone the Repository**:
    ```bash
    git clone https://github.com/n1tik91/To-Do-List-.git
    cd To-Do-List-
    ```

2. **Run the Program**:
    ```bash
    go run main.go
    ```

3. **Play the Game**:
    - Choose from the menu options:
      1. Add a task
      2. View tasks
      3. Remove a task
      4. Exit the program

---

## 🧩 Example Output

```plaintext
Welcome to the To-Do List Application!

Please choose an option:
1. Add a task
2. View tasks
3. Remove a task
4. Exit
1
Enter the task: Buy groceries
Task added!

Please choose an option:
1. Add a task
2. View tasks
3. Remove a task
4. Exit
2

Your tasks:
1. Buy groceries

Please choose an option:
1. Add a task
2. View tasks
3. Remove a task
4. Exit
3
Enter the task number to remove:
1
Task removed!

Please choose an option:
1. Add a task
2. View tasks
3. Remove a task
4. Exit
4
Exiting the To-Do List Application. Goodbye!
