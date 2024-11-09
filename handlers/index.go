package handlers

import (
	"fmt"
	"todo/controllers"
	"todo/data"
	models "todo/models"
)

func Insert() {
	var todo models.Todo

	controllers.Insert(todo)

	fmt.Println("all todos from handlers:", data.Todos)
}
