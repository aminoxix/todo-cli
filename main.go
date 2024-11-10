package main

import (
	// "log"
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"todo/cmd/handlers"
	"todo/shared/data"
	// "todo/shared/models"
	// "github.com/gorilla/mux"
)


func main() {

	// router := mux.NewRouter()
	// router.HandleFunc("/", serveHome)
	// log.Fatal(http.ListenAndServe(":8080", router))

	// var todo models.Todo
	// var id int

	fmt.Printf(`

What operation do you want to perform?

c: Insert a todo
r: View all todos
u: Update a todo
d: Delete a todo

`)

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("no input provided!")
	}
	formattedInput := strings.TrimSpace(input)

	switch formattedInput {
	case "c":
		handlers.Insert()
	case "r":
		handlers.ViewAll()
	case "u":
		handlers.Update()
	case "d":
		handlers.Delete()
	default:
		handlers.ViewAll()
	}

	// init
	data.Todos = handlers.ViewAll()

	// create file
	file, err := os.Create("./shared/data/todos.json")
	if err != nil {
		panic(err)
	}
	// for indentation
	finalJson, err := json.MarshalIndent(data.Todos, "", "\t")
	if err != nil {
		panic(err)
	}
	// write file
	os.WriteFile("./shared/data/todos.json", finalJson, 0644)
	// file close
	file.Close()

	fmt.Println("writing todos file...", data.Todos)
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>hey, this is a todo app</h1>"))
} 
