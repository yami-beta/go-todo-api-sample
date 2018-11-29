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
	Text     string `json:"text"`
	Complete bool   `json:"complete"`
}

// Response : API Response
type Response struct {
	Error   bool        `json:"error"`
	Payload interface{} `json:"payload"`
}

var (
	todos = map[int]Todo{
		1: Todo{Text: "todo1", Complete: false},
		2: Todo{Text: "todo2", Complete: false},
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

	todo := todos[id]
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(Response{
		Error:   false,
		Payload: todo,
	})
}

func createTodoHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	id := len(todos) + 1
	var todo Todo
	if err := decoder.Decode(&todo); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		json.NewEncoder(w).Encode(Response{Error: true, Payload: err.Error()})
		return
	}

	todos[id] = todo
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(Response{
		Error:   false,
		Payload: nil,
	})
}

func editTodoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		json.NewEncoder(w).Encode(Response{Error: true, Payload: err.Error()})
		return
	}
	todo := todos[id]

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&todo); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		json.NewEncoder(w).Encode(Response{Error: true, Payload: err.Error()})
		return
	}
	todos[id] = todo

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(Response{Error: false, Payload: nil})
}

func deleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		json.NewEncoder(w).Encode(Response{Error: true, Payload: err.Error()})
		return
	}

	delete(todos, id)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(Response{Error: false, Payload: nil})
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/todos", listTodoHandler).Methods("GET")
	r.HandleFunc("/todos", createTodoHandler).Methods("POST")
	r.HandleFunc("/todos/{id}", showTodoHandler).Methods("GET")
	r.HandleFunc("/todos/{id}", editTodoHandler).Methods("PATCH")
	r.HandleFunc("/todos/{id}", deleteTodoHandler).Methods("DELETE")
	// 本当はgraceful shutdown処理を書くべきだが省略
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
