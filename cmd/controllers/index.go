package controllers

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"todo/shared/data"
	models "todo/shared/models"
	utils "todo/shared/utils"
)

// insert structure
type insertModel struct {
	text     string
	function func(todo models.Todo) models.Todo
	continued bool
}

// update structure
type updateModel struct {
    text      string
    function  func() []models.Todo
    continued bool
}


// C for create
func Insert() models.Todo {
	var todo models.Todo
	crazyModels := []insertModel{
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
		return Insert()
	} else {
		return todo
	}
}

// R for read all
func ViewAll()[]models.Todo {
	var todos []models.Todo
	readJson, err := os.ReadFile("./shared/data/todos.json")
	if err != nil {
		panic(err)
	}
	checkValid := json.Valid(readJson)
	if checkValid {
		json.Unmarshal(readJson, &todos)
	}
	jsonData, err := json.MarshalIndent(todos, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", jsonData)
	return todos
}

// U for update
func Update() {
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


// utils begins here...
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

// for insert
func saveTodo(todo models.Todo) models.Todo {
	var isFileExist bool
	var todos []models.Todo

	todo.ID = int(utils.Utils().Int64())

	_, err := os.Stat("./shared/data/todos.json")
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
		file, err := os.Create("./shared/data/todos.json")
		if err != nil {
			panic(err)
		}
		file.Close()
	}

	data.Todos = append(todos, todo)
	fmt.Println("todos", data.Todos)
	return todo
}

// for update
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
// utils ends here
