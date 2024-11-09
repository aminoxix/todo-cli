package handlers

import (
	"todo/controllers"
	models "todo/models"
)

func Insert(todo models.Todo) {
	controllers.Insert(todo)
}

func ViewAll() {
	controllers.ViewAll()
}

func Update(id int) {
	controllers.Update(id)
}

func Delete() {
	controllers.Delete()
}
