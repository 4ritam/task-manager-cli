/*
Copyright © 2024 Pritam Das <daspritam4231@gmail.com>
*/

package core

type Task struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

func NewTask(id int, description string) *Task {
	return &Task{
		ID:          id,
		Description: description,
		Status:      "pending",
	}
}

func (t *Task) Complete() {
	t.Status = "completed"
}

func (t *Task) Incomplete() {
	t.Status = "pending"
}
