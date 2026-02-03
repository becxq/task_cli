package main

import (
	"encoding/json"
	"fmt"
	"os"
	"slices"
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
	Tasks []Task `json:"tasks"`
}

type Task struct { // отдельный тип для задач
	Description string    `json:"description"`
	Status      Status    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func (t *TaskLister) decode() {
	read, _ := os.ReadFile("tasks.json")
	json.Unmarshal(read, t)
}

func (t *TaskLister) encode() {
	data, _ := json.Marshal(t)

	os.WriteFile("tasks.json", data, 0644)
}

func (t *TaskLister) Add(description string) {
	task := Task{Description: description, Status: ToDo, CreatedAt: time.Now(), UpdatedAt: time.Now()}

	t.decode()
	t.Tasks = append(t.Tasks, task)
	t.encode()
}

func (t *TaskLister) Remove(index int) {
	t.decode()
	if index < 0 || index >= len(t.Tasks) {
		return
	}

	t.Tasks = slices.Delete(t.Tasks, index, index+1)
	t.encode()
}

func (t TaskLister) Show() {
	t.decode()
	for _, r := range t.Tasks {
		fmt.Printf("%s: %s | %s - %s", r.Description, r.Status, r.CreatedAt, r.UpdatedAt)
	}
}

func (t TaskLister) ShowByStatus(status Status) {
	t.decode()
	for _, r := range t.Tasks {
		if r.Status == status {
			fmt.Printf("%s: %s | %s - %s", r.Description, r.Status, r.CreatedAt, r.UpdatedAt)
		}
	}
}

func (t *TaskLister) UpdateStatus(status Status, index int) {
	t.decode()
	t.Tasks[index].Status = status
	t.encode()
}

func (t *TaskLister) UpdateName(description string, index int) {
	t.decode()
	t.Tasks[index].Description = description
	t.encode()
}
