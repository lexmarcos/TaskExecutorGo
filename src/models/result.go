package models

import "time"

type Result struct {
	ID     int
	Result int
	Time   time.Duration
}

func NewResult(id int, result int, time time.Duration) *Result {
	return &Result{
		ID:     id,
		Result: result,
		Time:   time,
	}
}
