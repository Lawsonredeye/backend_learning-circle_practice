# Simple Task Manager CLI application using Golang

This is a simple CLI application thats built for practice purpose with further integration and continuous iterations.

## Usage
It just creates a task on the command line when you run basic commands like **add**, and you can see all the tasks that has been created and stored by running the **list command** and to delete you run the delete command together with the task ID.

### Adding a task
```golang
go run main.go add "Your Task Name"
```
### Listing all Stored Tasks
```go
go run main.go list
```

### Delete a task
```go
go run main.go delete "Task ID"
```
