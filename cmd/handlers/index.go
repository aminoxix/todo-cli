package handlers

import (
	"todo/cmd/controllers"
	"todo/shared/models"
	// models "todo/shared/models"
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
