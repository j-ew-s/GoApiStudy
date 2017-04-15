package main

import (
	"time"
)

var id int

var todos Todos

func init() {
	Create(Todo{Name: "Create Kanban board", Completed: true, Due: time.Now().Add(1)})
	Create(Todo{Name: "Creat tasks cards"})
}

func FindByid(id int) Todo {
	var t Todo
	for _, t = range todos {
		if t.ID == id {
			return t
		}
	}
	return Todo{}
}

func Create(t Todo) Todo {
	id += 1
	t.ID = id
	todos = append(todos, t)
	return t
}
