# Requirements

Everyone can create todo lists and todo tasks.

## Domain

* ToDoList
  * Board can find by name in URI todo.example.com/todolist_name
  * If a new board with same name will be created the board will be overridden
  * If a ToDoList wasnt edit after, a day, a month or a year (can be choosed) it will be deleted
  * Share as readonly

* Task
  * Every task has an color, the color will mapped in frontend to a username
  * Tasks can have under tasks
  * Tasks can have a due date (the due date will shown in ui with colors)
  * Task can have a comment
  * Tasks can be ordered
