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

	for {
		/*
		*
		* Get Prompt case and check with switch
		*
		**/

		fmt.Println("enter your entry : ")
		res := readPrompt()

		switch res[0] {
		case "add":
			ID++
			fmt.Println(res)
			task.Add(ID, res[1])
		case "list":
			task.ShowAll()
		case "update":
			id, _ := strconv.Atoi(res[1])
			task.Update(id, res[2])
		}
	}
}
