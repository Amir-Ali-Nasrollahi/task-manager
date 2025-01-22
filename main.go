package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readPrompt() string {

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	text = strings.ReplaceAll(text, "\n", "")

	return text
}

func main() {

	task := Task{}

	value := readPrompt()
	fmt.Println(value)
	task.Add(2, "makeGoLang", "in-progcessing")
	task.ShowAll()
}
