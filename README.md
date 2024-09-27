A simple task list for the console.
Play around with cobra, gorm and sqlite.

# Usage
Build the app by calling `go build -o tasks .`.
This will compile tha app into an executable named `tasks` in current folder.
Launching `tasks` without any additional parameter, you will see the usage guide.

## Adding a new task
Add a new task to the todo list with the `create` command. 
Provide the text of the task right after create.

With the optional `due` flag, you can specify a due date. Specify the due date in a 
human readable format like `tomorrow`, `in 2 days`, `next week` or `december 23rd`.
```
./tasks create [<task>] [--due <due-date>]
```

### Example
```
./tasks create "Without due date"
./tasks create "Learn cobra" --due "next week"
```

## Set a task complete
Set a task as being complete. Specify the task to complete by it's id.
```
./tasks complete [id]
```
### Example
```
./tasks complete WPtjZvKso6Xq-ViC2_HzB
```

## Remove a task
Removes a task from the todo list by it's id.
```
./tasks delete [id]
```
### Example
```
./tasks delete WPtjZvKso6Xq-ViC2_HzB
```

## List all tasks
Lists all of your tasks in the todo list.
```
./tasks list
```
