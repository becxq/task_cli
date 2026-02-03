package main

import (
	"fmt"
	"os"
	"strconv"
)

func usage() {
	fmt.Println(`Usage:
  task-cli add "Description"
  task-cli update <id> "New description"
  task-cli delete <id>
  task-cli mark-in-progress <id>
  task-cli mark-done <id>
  task-cli list
  task-cli list <todo|in-progress|done>

Notes:
  <id> is 1-based (the first task is 1)

Examples:
  task-cli add "Buy groceries"
  task-cli update 1 "Buy groceries and cook dinner"
  task-cli delete 1
  task-cli mark-in-progress 1
  task-cli mark-done 1
  task-cli list
  task-cli list done
`)
}

func parseIDToIndex(s string) (int, error) {
	id, err := strconv.Atoi(s)
	if err != nil || id <= 0 {
		return 0, fmt.Errorf("invalid id: %q", s)
	}
	return id - 1, nil // превращаем "ID" (1..n) в index (0..n-1)
}

func parseStatusCLI(s string) (Status, error) {
	switch s {
	case "todo":
		return ToDo, nil
	case "in-progress":
		// пользователь вводит in-progress, а у тебя InProgress="in_progress"
		return InProgress, nil
	case "done":
		return Done, nil
	default:
		return "", fmt.Errorf("unknown status: %q (use todo|in-progress|done)", s)
	}
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		usage()
		os.Exit(1)
	}

	var tl TaskLister
	cmd := args[0]

	switch cmd {
	case "add":
		if len(args) < 2 {
			fmt.Fprintln(os.Stderr, "add requires description")
			usage()
			os.Exit(1)
		}
		tl.Add(args[1])
		fmt.Println("Task added successfully")

	case "update":
		if len(args) < 3 {
			fmt.Fprintln(os.Stderr, "update requires <id> and new description")
			usage()
			os.Exit(1)
		}
		index, err := parseIDToIndex(args[1])
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err)
			os.Exit(1)
		}
		tl.UpdateName(args[2], index)
		fmt.Println("Task updated successfully")

	case "delete":
		if len(args) < 2 {
			fmt.Fprintln(os.Stderr, "delete requires <id>")
			usage()
			os.Exit(1)
		}
		index, err := parseIDToIndex(args[1])
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err)
			os.Exit(1)
		}
		tl.Remove(index)
		fmt.Println("Task deleted successfully")

	case "mark-in-progress":
		if len(args) < 2 {
			fmt.Fprintln(os.Stderr, "mark-in-progress requires <id>")
			usage()
			os.Exit(1)
		}
		index, err := parseIDToIndex(args[1])
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err)
			os.Exit(1)
		}
		tl.UpdateStatus(InProgress, index)
		fmt.Println("Task marked as in progress")

	case "mark-done":
		if len(args) < 2 {
			fmt.Fprintln(os.Stderr, "mark-done requires <id>")
			usage()
			os.Exit(1)
		}
		index, err := parseIDToIndex(args[1])
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err)
			os.Exit(1)
		}
		tl.UpdateStatus(Done, index)
		fmt.Println("Task marked as done")

	case "list":
		// list или list <status>
		if len(args) == 1 {
			tl.Show()
			return
		}
		st, err := parseStatusCLI(args[1])
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err)
			os.Exit(1)
		}
		tl.ShowByStatus(st)

	default:
		fmt.Fprintln(os.Stderr, "Unknown command:", cmd)
		usage()
		os.Exit(1)
	}
}
