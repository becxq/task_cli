package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	var taskmanager TaskLister

	data, _ := os.ReadFile("tasks.json")
	json.Unmarshal(data, &taskmanager)

	fmt.Println(taskmanager.Count)
}
