package main

import (
	"fmt"
	"github.com/laziness/tasks"
	"bufio"
	"os"
)

func main() {
	manager := tasks.NewTaskManager()

	fmt.Println("Welcome! Work, Work Work!")

	for(true) {
		reader := bufio.NewReader(os.Stdin)

		fmt.Print("> ")
		command, _ := reader.ReadString('\n')
		command = command[:len(command)-1]

		parseCommand(command, manager)
	}
}
