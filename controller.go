package main

import (
	"time"
)

type Status string

const ( // enum для статуса
	ToDo       Status = "todo"
	InProgress Status = "in_progress"
	Done       Status = "done"
)

type TaskLister struct { // позволяет удобно взаимодействовать с tasks.json
	Count int    `json:"count"`
	Tasks []Task `json:"tasks`
}

type Task struct { // отдельный тип для задач
	Description string    `json:"description"`
	Status      Status    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
