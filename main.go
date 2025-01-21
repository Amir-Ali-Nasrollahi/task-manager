package main


func main() {

	task := Task{}

	task.Add(1, "read book", "TODO")

	task.Add(2, "make go lang", "in-progcessing")
	task.ShowAll()
}
