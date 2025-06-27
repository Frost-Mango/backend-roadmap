package main

import (
	"encoding/json"
	"io"
	"os"
)

func readObject(fp string, n int64) ([]byte, int64, error) {
	f, err := os.Open(fp)
	if err == nil {
		var m int
		b := make([]byte, 1)
		obj := ""
		_, err := f.Seek(n, io.SeekStart)
		if err != nil {
			return make([]byte, 0), -1, err
		}
		for {
			m, err = f.Read(b)
			if err != nil {
				return make([]byte, 0), -1, err
			}
			s := string(b)
			obj += s
			if s == "}" {
				return []byte(obj), int64(m), nil
			}
		}
	}
	return make([]byte, 0), -1, err
}

func importTasks(fp string) []*Task {
	var tasks []*Task
	var n int64 = 0
	for {
		obj, n, err := readObject(fp, n)
		if err != nil {
			break
		}
		var t *Task
		err = json.Unmarshal(obj, t)
		check(err)
		tasks = append(tasks, t)
		n += 1
	}
	return tasks
}
