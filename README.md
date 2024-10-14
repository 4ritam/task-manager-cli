# CLI Task Manager

A simple command-line interface (CLI) task manager written in Go.

## Features

- Add new tasks
- List all tasks
- Delete tasks
- Mark tasks as complete or incomplete
- Search for tasks

## Installation

If you have Make installed on your system, you can simply run:
```
make install
```

Otherwise, you can follow the steps below:

1. Ensure you have Go installed on your system.
2. Clone this repository:
   ```
   git clone https://github.com/yourusername/cli-task-manager.git
   ```
3. Navigate to the project directory:
   ```
   cd cli-task-manager
   ```
4. Build the project:
   ```
   go build
   ```

## Usage

### Add a task
```
./task-manager add "Buy milk"
```

### List all tasks
```
./task-manager list
```

### Delete a task
```
./task-manager delete 1
```

### Mark a task as complete
```
./task-manager complete 1
```

### Mark a task as incomplete
```
./task-manager incomplete 1
```

### Search for tasks
```
./task-manager search milk
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Author

Pritam Das - [daspritam4231@gmail.com](mailto:daspritam4231@gmail.com)
