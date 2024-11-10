package main

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2/log"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello, World!")

	r := mux.NewRouter()
	r.HandleFunc("/", WriteHome)
	r.HandleFunc("/{id}", Request)

	log.Fatal(http.ListenAndServe(":8080", r))
}

func WriteHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello, World!</h1>"))
}

func Request(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	vars := mux.Vars(r) // access all the variables passed in the URL path from the request
	w.Write([]byte(`{"message": "Request with id: ` + vars["id"] + `"}`))
	fmt.Println(vars)
}
