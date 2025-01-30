package main

import (
	"errors"
)

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

func (t Task) checkTodoExist(id int) (int, error) {
	var err error = errors.New("the id out of the range")
	var todo_key int
	for key, value := range t.todo {

		if value.ID <= id && value.ID == id {
			todo_key = key
			err = nil
			break
		}

	}

	return todo_key, err
}

func (t *Task) Add(ID int, Name string) {
	nextID := ID + 1

	t.todo = append(t.todo, Todo{ID, Name, "todo"})

	t.NextID = nextID

}

func (t Task) ShowAll() []Todo {
	return t.todo
}

func (t *Task) Update(id int, newName string) {

	todo_key, err := t.checkTodoExist(id)
	if err == nil {
		t.todo[todo_key].Name = newName
	}
}

func (t *Task) Delete(id int) {

	delete_key, err := t.checkTodoExist(id)
	if err == nil {
		t.todo = append(t.todo[0:delete_key], t.todo[delete_key+1:]...)
	}

}
func (t Task) ShowBy(showingBy string) []Todo {

	var list_of_todos []Todo
	for _, value := range t.todo {

		if value.Status == showingBy {
			list_of_todos = append(list_of_todos, value)
		}
	}
	return list_of_todos

}
func (t *Task) MarkStatus(id int, status string) {
	returnId, err := t.checkTodoExist(id)
	if err == nil {
		t.todo[returnId].Status = status
	}
}
