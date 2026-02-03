package main

import "time"

type Status string

const (
	ToDo       Status = "todo"
	InProgress Status = "in_progress"
	Done       Status = "done"
)

type Task struct {
	id          int
	description string
	status      Status
	createdAt   time.Time
	updatedAt   time.Time
}

func (t Task) UpdateStatus(status Status) {
	t.status = status
	t.updatedAt = time.Now()
}

func (t Task) Show() (string, Status) {
	return t.description, t.status
}

TaskController := make(map[Task.description]Task)