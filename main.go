package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Todo : Todo
type Todo struct {
	ID       int    `json:"id"`
	Text     string `json:"text"`
	Complete bool   `json:"complete"`
}

// Todos : slice of Todo
type Todos []Todo

// Response : API Response
type Response struct {
	Error   bool        `json:"error"`
	Payload interface{} `json:"payload"`
}

var (
	todos = Todos{
		Todo{ID: 1, Text: "todo1", Complete: false},
		Todo{ID: 2, Text: "todo2", Complete: false},
	}
)

func listTodoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(Response{
		Error:   false,
		Payload: todos,
	})
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/todos", listTodoHandler).Methods("GET")
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
