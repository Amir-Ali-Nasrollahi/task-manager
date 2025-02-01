package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readPrompt() []string {

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	text = strings.ReplaceAll(text, "\n", "")
	explode := strings.Split(text, " ")

	return explode
}

func main() {

	task := Task{}
	var ID int = 0


	fmt.Print("\n\nthe command is : `add`,`update`,`mark-`,`delete`\n")
	for {
		/*
		*
		* Get Prompt case and check with switch
		*
		**/

		fmt.Print("\nenter your entry : ")

		res := readPrompt()

		switch res[0] {
		case "add":
			ID++
			task.Add(ID, res[1])
			fmt.Printf("the Task added with ID : %v\n", ID)
		case "list":
			if len(res) == 1 {
				fmt.Println(task.ShowAll())
			} else {
				fmt.Println(task.ShowBy(res[1]))
			}
		case "update":
			id, _ := strconv.Atoi(res[1])
			task.Update(id, res[2])
			fmt.Printf("the Task updated with ID : %v\n", id)
		case "delete":
			id, _ := strconv.Atoi(res[1])
			task.Delete(id)
			fmt.Printf("the Task deleted with ID : %v\n", id)
		case "mark-done":
			id, _ := strconv.Atoi(res[1])
			task.MarkStatus(id, "done")
		case "mark-in-progress":
			id, _ := strconv.Atoi(res[1])
			task.MarkStatus(id, "in-progress")
		}
	}
}
