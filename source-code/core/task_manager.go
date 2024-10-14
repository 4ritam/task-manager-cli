/*
Copyright © 2024 Pritam Das <daspritam4231@gmail.com>
*/

package core

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type TaskManager struct {
	Tasks  []*Task `json:"tasks"`
	LastID int     `json:"lastID"`
}

func NewTaskManager() *TaskManager {
	return &TaskManager{
		Tasks:  []*Task{},
		LastID: 0,
	}
}

func (tm *TaskManager) AddTask(description string) {
	tm.loadState()
	defer tm.saveState()
	tm.LastID++
	t := NewTask(tm.LastID, description)
	tm.Tasks = append(tm.Tasks, t)
}

func (tm *TaskManager) GetTasks() []*Task {
	tm.loadState()
	return tm.Tasks
}

func (tm *TaskManager) getTask(id int) *Task {
	tm.loadState()
	for _, t := range tm.Tasks {
		if t.ID == id {
			return t
		}
	}
	return nil
}

func (tm *TaskManager) CompleteTask(id int) {
	tm.loadState()
	defer tm.saveState()
	t := tm.getTask(id)
	if t != nil {
		t.Complete()
	}
}

func (tm *TaskManager) IncompleteTask(id int) {
	tm.loadState()
	defer tm.saveState()
	t := tm.getTask(id)
	if t != nil {
		t.Incomplete()
	}
}

func (tm *TaskManager) DeleteTask(id int) {
	tm.loadState()
	defer tm.saveState()
	for i, t := range tm.Tasks {
		if t.ID == id {
			tm.Tasks = append(tm.Tasks[:i], tm.Tasks[i+1:]...)
			break
		}
	}
}

func (tm *TaskManager) SearchTasks(query string) []*Task {
	tm.loadState()
	var result []*Task
	for _, t := range tm.Tasks {

		if strings.Contains(strings.ToLower(t.Description), strings.ToLower(query)) {
			result = append(result, t)
		}
	}
	return result
}

func (tm *TaskManager) saveState() {
	// Save the state of the task manager to a file
	jsonData, err := json.MarshalIndent(tm, "", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = os.WriteFile("data/data.json", jsonData, 0644)

	if err != nil {
		fmt.Println(err)
	}

}

func (tm *TaskManager) loadState() {
	// Load the state of the task manager from a file
	jsonData, err := os.ReadFile("data/data.json")
	if err != nil {
		return
	}
	err = json.Unmarshal(jsonData, tm)
	if err != nil {
		fmt.Println(err)
	}
}
