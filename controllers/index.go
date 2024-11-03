package controllers

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	appTypes "todo/types"
	appUtils "todo/utils"
)

type model[] struct {
	text string
	function func(todo appTypes.Todo) appTypes.Todo
	continued bool
}

type completedModel struct {
    choices  []string           // items on the to-do list
    cursor   int                // which to-do list item our cursor is pointing at
    selected map[int]struct{}   // which to-do items are selected
}

var task string
var checked bool

func Insert(todo appTypes.Todo) appTypes.Todo {
	crazyModels := model{
		{
			text: "Enter your task?",
			function: inputTask,
			continued: true,
		},
		{
			text: "Is it completed?",
			function: checkCompletion,
			continued: true,
		},
		{
			text: "Thanks! Your response follows:",
			function: Insert,
			continued: false,
		},
	}

	for _, value := range crazyModels {
		if value.continued {
			fmt.Println(value.text)
			value.function(todo)
			// fmt.Printf("%d:\n text: %v\n function: %v\n continued: %v\n", key, value.text, value.function(), value.continued)
			continue
		}
	}

	// fmt.Println(appUtils.Utils())

	// s := "Enter your task?"
	// fmt.Println(s)
	// reader := bufio.NewReader(os.Stdin)
	// input, err := reader.ReadString('\n')
	// if err != nil {
	// 	log.Fatal("no input provided!")
	// }
	// task := strings.TrimSpace(input)

	// m := completedModel{
	// 	choices: []string{"yes", "no"},
	// 	selected: make(map[int]struct{}),
	// }

	// var isChecked bool

	// s = "Is it completed?"
	// fmt.Println(s)
	// // Iterate over our choices
    // for i, choice := range m.choices {

    //     // Is the cursor pointing at this choice?
    //     cursor := " " // no cursor
    //     if m.cursor == i {
    //         cursor = ">" // cursor!
    //     }

    //     // Is this choice selected?
    //     checked := " " // not selected
    //     if _, ok := m.selected[i]; ok {
    //         checked = "x" // selected!
    //     }

	// 	isChecked = checked == "yes"

    //     // Render the row
    //     s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
    // }

    // // The footer
    // s += "\nPress q to quit.\n"

    // // Send the UI for rendering

	todo.ID = int(appUtils.Utils().Int64())
	todo.Task = task
	todo.Checked = checked

	fmt.Println("todo from ctrls", todo)
	return todo
}

func inputTask(todo appTypes.Todo) appTypes.Todo {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("no input provided!")
	}
	task = strings.TrimSpace(input)

	return appTypes.Todo{
		Task: task,
	}
}

func checkCompletion(todo appTypes.Todo) appTypes.Todo {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	
	if err != nil {
		log.Fatal("no input provided!")
	}

	checkCheck := strings.TrimSpace(input)

	switch checkCheck {
		case "yes", "y":
			checked = true
		case "no", "n":
			checked = false
		default:
			checked = true
	}

	return appTypes.Todo{
		Checked: checked,
	}
}

// func Delete(todoId int) {

// }
