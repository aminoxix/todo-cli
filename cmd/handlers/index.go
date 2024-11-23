package handlers

import (
	"github.com/aminoxix/todo-cli/cmd/controllers"
	"github.com/aminoxix/todo-cli/shared/models"
	// models "github.com/aminoxix/todo-cli/shared/models"
)

func Insert() models.Todo {
	return controllers.Insert()
}

func ViewAll() []models.Todo {
	return controllers.ViewAll()
}

func Update() {
	controllers.Update()
}

func Delete() {
	controllers.Delete()
}
