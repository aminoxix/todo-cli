package controllers

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"todo/data"
	models "todo/models"
	utils "todo/utils"
)

type model struct {
	text     string
	function func(todo models.Todo) models.Todo
	continued bool
}

// C for create
func Insert(todo models.Todo) models.Todo {
	crazyModels := []model{
		{
			text:     "Enter your task?",
			function: inputTask,
			continued: true,
		},
		{
			text:     "Is it completed? (yes/no) ",
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
	var isFileExist bool
	var todos []models.Todo

	todo.ID = int(utils.Utils().Int64())

	_, err := os.Stat("./data/todos.json")
	if os.IsNotExist(err) {
    	isFileExist = false
	} else {
		isFileExist = true
	}

	if isFileExist {
		if len(data.Todos) != 0 {
			todos = data.Todos
		} else {
			todos = ViewAll()
		}
	} else {
		file, err := os.Create("./data/todos.json")
		if err != nil {
			panic(err)
		}
		file.Close()
	}

	data.Todos = append(todos, todo)
	fmt.Println("todos", data.Todos)
	return todo
}

// R for read all
func ViewAll()[]models.Todo {
	var todos []models.Todo
	readJson, err := os.ReadFile("./data/todos.json")
	if err != nil {
		panic(err)
	}
	checkValid := json.Valid(readJson)
	if checkValid {
		json.Unmarshal(readJson, &todos)
	}
	fmt.Println(todos)
	return todos
}

type updateModel struct {
    text      string
    function  func() []models.Todo
    continued bool
}

// U for update
func Update(id int) {
    crazyModels := []updateModel{
        {
            text: "Here's your todos:",
            function: func() []models.Todo {
                return ViewAll()
            },
            continued: true,
        },
        {
            text: "Enter ID to edit:",
            function: updateTodo,
            continued: true,
        },
        {
            text: "Thanks for update!",
            function: func() []models.Todo {
                return ViewAll()
            },
            continued: false,
        },
    }

    for _, value := range crazyModels {
        fmt.Println(value.text)
        value.function()

        if !value.continued {
            break
        }
    }
}

func updateTodo() []models.Todo {
    reader := bufio.NewReader(os.Stdin)
    idInput, err := reader.ReadString('\n')
    if err != nil {
        panic(err)
    }
    idString := strings.TrimSpace(idInput)
	id, err := strconv.Atoi(idString)
	if err != nil {
		panic(err)
	}

	fmt.Println("Enter updated task title:")
	taskInput, err := reader.ReadString('\n')
	if err != nil {
			panic(err)
		}
	task := strings.TrimSpace(taskInput)

	data.Todos = ViewAll()

    for i := range data.Todos {
        if data.Todos[i].ID == id {
            data.Todos[i].Task = task
            break
        }
    }

    fmt.Println("todo updated successfully!")
    return data.Todos
}

// D for delete
func Delete() {
	data.Todos = ViewAll()
	fmt.Println("Here's your todos:", data.Todos)

	fmt.Println("Enter todo ID you want to remove:")
	reader := bufio.NewReader(os.Stdin)
    idInput, err := reader.ReadString('\n')
    if err != nil {
        panic(err)
    }
    idString := strings.TrimSpace(idInput)
	id, err := strconv.Atoi(idString)
	if err != nil {
		panic(err)
	}

	for i := range data.Todos {
		if data.Todos[i].ID == id {
			data.Todos = append(data.Todos[:i], data.Todos[i+1:]...)
			fmt.Println("todo removed successfully!")
		}
	}
}
