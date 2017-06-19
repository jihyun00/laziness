package main

import (
	"fmt"
	"strconv"
	"strings"
	"os"

	"github.com/laziness/tasks"
)

func parseCommand(command string, manager *tasks.TaskManager) {
	commands := strings.SplitAfterN(command, " ", 2)

	op := strings.Trim(commands[0], " ")

	switch op {
	case "add":
		tasks.CreateTask(commands[1], manager)

		fmt.Println("Add task successfully")

	case "done":
		id, _ := strconv.Atoi(commands[1])
		tasks.CloseTask(id, manager)

		fmt.Println(id, "is checked done")

	case "exit":
		fmt.Println("Bye Bye")
		os.Exit(0)

	case "help":
		fmt.Println("Commands: \n"+
						"add \t\t Add your task.\n" +
						"\t\t e.g) add read a book\n" +
						"done \t\t Check done your task\n" +
						"\t\t e.g) done 2\n" +
						"exit \t\t Finish todo list.\n" +
			    		"\t\t e.g) exit\n" +
						"help \t\t Help about any command.\n" +
						"\t\t e.g) help\n" +
						"list \t\t Print your todo lists\n" +
						"\t\t e.g) list")

	case "list":
		tasks.PrintTasks(manager)
	}
}
