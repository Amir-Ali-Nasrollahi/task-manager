package main

import "fmt"

type Todo struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}
type Task struct {
	/*
	
	* it means make slice which each item be a Todo struct
	* 
	* if you say []int -> means make slice that every item is int

	**/
	todo   []Todo
	NextID int
}

func (t *Task) Add(ID int, Name string, Status string) {
	nextID := ID + 1

	t.todo = append(t.todo, Todo{ID, Name, Status})

	t.NextID = nextID

}

func (t Task) ShowAll() {

	fmt.Println(t.todo)

}
