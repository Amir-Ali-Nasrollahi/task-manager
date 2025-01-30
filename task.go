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

func (t Task) CheckTodoExist(id int) (int, int) {
	flag := 0
	var todo_key int
	for key, value := range t.todo {

		if value.ID <= id && value.ID == id {
			todo_key = key
			flag = 1
			break
		}

	}

	return todo_key, flag
}

func (t *Task) Add(ID int, Name string) {
	nextID := ID + 1

	t.todo = append(t.todo, Todo{ID, Name, "todo"})

	t.NextID = nextID

}

func (t Task) ShowAll() {

	fmt.Println(t.todo)

}

func (t *Task) Update(id int, newName string) {

	flag := 0
	var todo_key int
	for key, value := range t.todo {

		if value.ID <= id && value.ID == id {
			todo_key = key
			flag = 1
			break
		}

	}
	if flag == 1 {
		t.todo[todo_key] = Todo{id, newName, "todo"}
	}
}

func (t *Task) Delete(id int) {

}
