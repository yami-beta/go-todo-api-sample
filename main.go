package main

import (
	"encoding/json"
	"html"
	"log"
	"net/http"
)

// Response : API Response
type Response struct {
	Message string `json:"message"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	res := Response{Message: "Hello, " + html.EscapeString(r.URL.Path)}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(res)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
