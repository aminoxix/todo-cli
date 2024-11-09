package controllers

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"todo/data"
	models "todo/models"
	appUtils "todo/utils"
)

type model struct {
	text     string
	function func(todo models.Todo) models.Todo
	continued bool
}

func Insert(todo models.Todo) models.Todo {
	crazyModels := []model{
		{
			text:     "Enter your task?",
			function: inputTask,
			continued: true,
		},
		{
			text:     "Is it completed?",
			function: checkCompletion,
			continued: true,
		},
		{
			text:     "Thanks for your response!",
			function: saveTodo,
			continued: false,
		},
	}

	var choice string
	for _, value := range crazyModels {
		fmt.Println(value.text)
		todo = value.function(todo)

		if !value.continued {
			fmt.Printf("Do you want to continue? (yes/no) ")
			reader := bufio.NewReader(os.Stdin)
			input, err := reader.ReadString('\n')
			if err != nil {
				log.Fatal("no input provided!")
			}
			choice = strings.TrimSpace(input)
			break
		}
	}

	if choice == "yes" || choice == "y" {
		return Insert(todo)
	} else {
		return todo
	}
}

func inputTask(todo models.Todo) models.Todo {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("no input provided!")
	}
	todo.Task = strings.TrimSpace(input)
	return todo
}

func checkCompletion(todo models.Todo) models.Todo {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("no input provided!")
	}

	checkCheck := strings.TrimSpace(input)
	switch checkCheck {
	case "yes", "y":
		todo.Checked = true
	case "no", "n":
		todo.Checked = false
	default:
		todo.Checked = true
	}

	return todo
}

func saveTodo(todo models.Todo) models.Todo {
	todo.ID = int(appUtils.Utils().Int64())
	data.Todos = append(data.Todos, todo)
	return todo
}