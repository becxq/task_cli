package main

import (
	"time"
)

type Status string

const (
	ToDo       Status = "todo"
	InProgress Status = "in_progress"
	Done       Status = "done"
)

type TaskLister struct {
	Count int    `json:"count"`
	Tasks []Task `json: "tasks`
}

type Task struct {
	Id          int       `json:"tasks"`
	Description string    `json:"tasks"`
	Status      Status    `json:"tasks"`
	CreatedAt   time.Time `json:"tasks"`
	UpdatedAt   time.Time `json:"tasks"`
}

func (t Task) UpdateStatus(status Status) {

}

func (t Task) Show() (string, Status) {

}

func (t Task) Delete() {

}

func (t TaskLister) Create(description string, status Status) {

}
