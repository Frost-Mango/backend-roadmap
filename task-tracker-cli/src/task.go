package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Task struct {
	Id        int
	Desc      string
	Status    string
	CreatedAt string
	UpdatedAt string
}

func createTask(id int, desc string) *Task {
	return &Task{
		Id:        id,
		Desc:      desc,
		Status:    "todo",
		CreatedAt: time.Now().Format(time.UnixDate),
		UpdatedAt: time.Now().Format(time.UnixDate),
	}
}

func updateTask(t *Task, newDesc string) {
	t.Desc = newDesc
	t.UpdatedAt = time.Now().Format(time.UnixDate)
}

func markTask(t *Task, st string) {
	t.Status = st
	t.UpdatedAt = time.Now().Format(time.UnixDate)
}

func checkStatus(t *Task, ts string) bool {
	return t.Status == ts
}

func printTask(t *Task) {
	fmt.Println("------------")
	fmt.Println(t.Id, t.Desc)
	fmt.Println(t.Status)
	fmt.Println("created at:", t.CreatedAt)
	fmt.Println("updated at:", t.UpdatedAt)
}

func export(tasks []*Task) {
	res, err := json.Marshal(tasks)
	check(err)
	err = os.WriteFile("./tasks.json", res, 0644)
}
