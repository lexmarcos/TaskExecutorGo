package models

type Task struct {
	ID    int
	Cost  float64
	Type  string
	Value int
}

func NewTask(id int, cost float64, taskType string, value int) *Task {
	return &Task{
		ID:    id,
		Cost:  cost,
		Type:  taskType,
		Value: value,
	}
}
