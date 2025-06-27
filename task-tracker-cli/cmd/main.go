package main

import (
	"encoding/json"
	"fmt"
	"os"
	"slices"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	tasks := importTasks()
	export(tasks)
}

func importTasks() []*Task {
	var tasks []*Task
	f, err := os.Open("./tasks.json")
	if err == nil {
		b := make([]byte, 1)
		in_obj := false
		co := ""
		for {
			_, err = f.Read(b)
			if err != nil {
				return tasks
			}
			s := string(b)
			switch s {
			case "n":
				if !in_obj {
					return tasks
				} else {
					co += s
				}
			case "{":
				in_obj = true
				co += s
			case "}":
				in_obj = false
				co += s
				var t Task
				err = json.Unmarshal([]byte(co), &t)
				tasks = append(tasks, &t)
				co = ""
			default:
				if in_obj {
					co += s
				}
			}
		}
	}
	os.Create("./tasks.json")
	return tasks
}

func addTask(tasks []*Task, desc string) []*Task {
	id := 0
	if len(tasks) != 0 {
		id = int(tasks[len(tasks)-1].Id) + 1
	}
	return append(tasks, createTask(id, desc))
}

func getTaskInd(tasks []*Task, id int) (int, error) {
	n, found := slices.BinarySearchFunc(tasks, &Task{id, "", "", "", ""}, func(a, b *Task) int {
		return a.Id - b.Id
	})
	if !found {
		return 0, fmt.Errorf("No task with provided ID") //TODO: write proper errors
	}
	return n, nil
}

func getTask(tasks []*Task, id int) (*Task, error) {
	n, err := getTaskInd(tasks, id)
	if err == nil {
		return tasks[n], nil
	}
	return &Task{id, "", "", "", ""}, err
}

func delete(tasks []*Task, id int) ([]*Task, error) {
	n, err := getTaskInd(tasks, id)
	if err == nil {
		return slices.Delete(tasks, n, n+1), nil
	}
	return tasks, err
}

func update(tasks []*Task, id int, newDesc string) error {
	t, err := getTask(tasks, id)
	if err == nil {
		updateTask(t, newDesc)
	}
	return err
}

func markInProgress(tasks []*Task, id int) error {
	t, err := getTask(tasks, id)
	if err == nil {
		markTask(t, "in progress")
	}
	return err
}

func markDone(tasks []*Task, id int) error {
	t, err := getTask(tasks, id)
	if err == nil {
		markTask(t, "done")
	}
	return err
}

func list(tasks []*Task) {
	for i := range len(tasks) {
		printTask(tasks[i])
	}
}

func listFiltered(tasks []*Task, filter string) error {
	switch filter {
	case "in progress", "done", "todo":
		for i := range len(tasks) {
			if tasks[i].Status == filter {
				printTask(tasks[i])
			}
		}
	default:
		return fmt.Errorf("Invalid filter") //TODO: write proper errors
	}
	return nil
}

func export(tasks []*Task) {
	res, err := json.Marshal(tasks)
	check(err)
	err = os.WriteFile("./tasks.json", res, 0644)
}
