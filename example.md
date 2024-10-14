# CLI Task Manager: Usage Example

This demonstration shows how a user might interact with the CLI Task Manager application. The `$` symbol represents the command prompt.

```
$ ./task-manager

Welcome to the Task Manager!
Available commands:
  add <task description>
  list
  complete <task id>
  delete <task id>
  search <keyword>
  exit

> add Buy groceries
Task added successfully. ID: 1

> add Finish Go project
Task added successfully. ID: 2

> add Call mom
Task added successfully. ID: 3

> list
Tasks:
1. [ ] Buy groceries
2. [ ] Finish Go project
3. [ ] Call mom

> complete 1
Task 1 marked as completed.

> list
Tasks:
1. [x] Buy groceries
2. [ ] Finish Go project
3. [ ] Call mom

> delete 3
Task 3 deleted successfully.

> list
Tasks:
1. [x] Buy groceries
2. [ ] Finish Go project

> search project
Matching tasks:
2. [ ] Finish Go project

> add Schedule dentist appointment
Task added successfully. ID: 3

> list
Tasks:
1. [x] Buy groceries
2. [ ] Finish Go project
3. [ ] Schedule dentist appointment

> exit
Tasks saved. Goodbye!

$
```

In this example:
- The application starts with a welcome message and lists available commands.
- Users can add tasks using the `add` command followed by the task description.
- The `list` command shows all tasks with their IDs and completion status.
- Users can mark tasks as complete using the `complete` command and the task ID.
- Tasks can be deleted using the `delete` command and the task ID.
- The `search` command allows users to find tasks containing specific keywords.
- The application automatically saves tasks between sessions.
- Users can exit the application using the `exit` command.

Note: This is a basic example. Your implementation could include additional features like priority levels, due dates, or more advanced search capabilities.
