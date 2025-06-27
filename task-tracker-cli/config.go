package main

import (
	"encoding/json"
	"os"
)

type Config struct {
	FilePath string
}

func readConfig() Config {
	obj, _, err := readObject("config.json", 0)
	if err != nil {
		return defaultConfig()
	}
	var cnf Config
	check(json.Unmarshal(obj, &cnf))
	return cnf
}

func defaultConfig() Config {
	f, err := os.Create("config.json")
	check(err)
	defer f.Close()
	cnf := Config{
		FilePath: "tasks.task",
	}
	res, err := json.Marshal(cnf)
	check(err)
	_, err = f.Write(res)
	check(err)
	return cnf
}
