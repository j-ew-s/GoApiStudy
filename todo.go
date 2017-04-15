package main

import "time"

/***Todo***/
type Todo struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

/***List of Todo***/
type Todos []Todo
