package task

import (
	"fmt"
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
