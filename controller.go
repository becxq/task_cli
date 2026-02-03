package main

import (
	"time" 
	"encoding/json"
)

type Status string

const (
	ToDo       Status = "todo"
	InProgress Status = "in_progress"
	Done       Status = "done"
)

type Task struct {
	id          int       `json:"id"`
	description string    `json:"description"`
	status      Status    `json:"status"`
	createdAt   time.Time `json:"createdAt"`
	updatedAt   time.Time `json:"updatedAt"`
}

func (t Task) UpdateStatus(status Status) {
	t.status = status
	t.updatedAt = time.Now()
}

func (t Task) Show() (string, Status) {
	return t.description, t.status
}

func (t Task) Delete() {
	delete(TaskController, t.id)
}

func Create(description string, status Status) {
	TaskController[] = 
}
