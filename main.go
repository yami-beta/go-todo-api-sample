package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

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

func showTodoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		json.NewEncoder(w).Encode(Response{Error: true, Payload: err.Error()})
		return
	}

	var todo *Todo
	for _, v := range todos {
		if v.ID == id {
			todo = &v
			break
		}
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(Response{
		Error:   false,
		Payload: todo,
	})
}

func createTodoHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	id := len(todos) + 1
	todo := Todo{ID: id}
	if err := decoder.Decode(&todo); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		json.NewEncoder(w).Encode(Response{Error: true, Payload: err.Error()})
		return
	}

	todos = append(todos, todo)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(Response{
		Error:   false,
		Payload: nil,
	})
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/todos", listTodoHandler).Methods("GET")
	r.HandleFunc("/todos", createTodoHandler).Methods("POST")
	r.HandleFunc("/todos/{id}", showTodoHandler).Methods("GET")
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
