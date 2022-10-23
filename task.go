package main

import (
	"encoding/json"
	"log"
	"os"
)

type status int

// iota is a shortcut for incrementing constants. starts at 0
// it is presumed that constants within a single declaration all have the same expression
// this is why the shortcut below works to set todo, inProgress and done as status types with values 0, 1 and 2
const (
	todo status = iota
	inProgress
	done
)

type Task struct {
	Status  status `json:"status"`
	Name    string `json:"title"`
	Details string `json:"description"`
}

// implement list.Item interface
func (t Task) FilterValue() string {
	return t.Name
}

func (t Task) Title() string {
	return t.Name
}

func (t Task) Description() string {
	return t.Details
}

func FromJSONFile(filename string) []Task {
	f, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal("Problem reading JSON file.", err)
	}
	var tasks []Task
	json.Unmarshal(f, &tasks)
	return tasks
}
