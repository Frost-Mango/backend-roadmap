package main

import (
	"encoding/json"
	"io"
	"os"
)

func readObject(fp string, n int64) ([]byte, int64, error) {
	f, err := os.Open(fp)
	if err != nil {
		return make([]byte, 0), -1, err
	}
	defer f.Close()
	var m int64 = n
	b := make([]byte, 1)
	obj := ""
	_, err = f.Seek(n, io.SeekStart)
	if err != nil {
		return make([]byte, 0), -1, err
	}
	for {
		_, err = f.Read(b)
		m += 1
		if err != nil {
			return make([]byte, 0), -1, err
		}
		s := string(b)
		obj += s
		if s == "}" {
			return []byte(obj), m, nil
		}
	}
}

func importTasks(fp string) []*Task {
	var tasks []*Task
	var n int64 = 0
	for {
		obj, m, err := readObject(fp, n)
		if err != nil {
			break
		}
		var t Task
		err = json.Unmarshal(obj, &t)
		check(err)
		tasks = append(tasks, &t)
		n = m
	}
	return tasks
}

func export(fp string, tasks []*Task) {
	f, err := os.Create(fp)
	check(err)
	defer f.Close()
	for i := range len(tasks) {
		res, err := json.Marshal(tasks[i])
		check(err)
		_, err = f.Write(res)
		check(err)
	}
	f.Sync()
}
