package main

import (
	// "log"
	"net/http"
	"todo/handlers"
	// "github.com/gorilla/mux"
)


func main() {

	// router := mux.NewRouter()
	// router.HandleFunc("/", serveHome)
	// log.Fatal(http.ListenAndServe(":8080", router))

	handlers.Insert()
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>hey, this is a todo app</h1>"))
} 
