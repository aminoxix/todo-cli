package handlers

import (
	"fmt"

	"todo/controllers"
	appData "todo/data"
	appTypes "todo/types"
)

func Insert() {
	var todo appTypes.Todo

	insertedTodo := controllers.Insert(todo)

	todos :=  append(appData.Todos, insertedTodo)

	// fmt.Println("all todos", todos)
	// fmt.Println("todo from handler", insertedTodo)

	fmt.Println("all", todos)
}
