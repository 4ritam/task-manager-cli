# Project 1: CLI Task Manager

## Objective
Create a command-line interface (CLI) task manager application in Go that allows users to manage their tasks.

## Features
1. Add new tasks
2. List all tasks
3. Mark tasks as completed
4. Delete tasks
5. Save tasks to a file
6. Load tasks from a file

## Implementation Steps

1. Set up the project:
   - Create a new Go module
   - Set up the main package and entry point

2. Define the task structure:
   - Create a struct to represent a task (id, description, status)

3. Implement task management functions:
   - Add task
   - List tasks
   - Mark task as completed
   - Delete task

4. Create the CLI interface:
   - Use the "flag" package or a third-party CLI library like "cobra"
   - Implement command-line arguments for each action

5. Implement data persistence:
   - Create functions to save tasks to a JSON file
   - Create functions to load tasks from a JSON file

6. Add error handling:
   - Implement proper error checking and handling throughout the application

7. Implement a simple REPL (Read-Eval-Print Loop):
   - Allow users to enter commands continuously without restarting the program

8. Add data validation:
   - Ensure task descriptions are not empty
   - Validate user inputs

9. Implement basic searching and filtering:
   - Allow users to search tasks by keyword
   - Add ability to filter tasks by status (completed/uncompleted)

10. Optimize and refactor:
    - Review your code for any repetitions or inefficiencies
    - Refactor to make the code more modular and maintainable

11. Add unit tests:
    - Write tests for core functions using the Go testing package

12. Document your code:
    - Add comments explaining complex parts of your code
    - Create a README file with instructions on how to use the application

## Bonus Challenges
- Implement task priority levels
- Add due dates to tasks and sorting by due date
- Create sub-tasks or task dependencies
- Implement data encryption for the saved tasks file

## Learning Outcomes
- Basic Go syntax and data structures
- Working with the Go standard library
- File I/O operations in Go
- JSON marshaling and unmarshaling
- Command-line application development
- Basic error handling in Go
- Unit testing in Go
