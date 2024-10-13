package dtos

import "time"

type Todo struct {
	ID         int
	Title      string
	Complete   bool
	CompleteBy time.Time
}

func NewTodo(id int, title string, complete bool, completeBy time.Time) *Todo {
	return &Todo{ID: id, Title: title, Complete: complete, CompleteBy: completeBy}
}
