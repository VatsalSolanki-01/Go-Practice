package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Task struct {
	ID    int
	Title string
	Done  bool
}

var tasks = []Task{
	{ID: 1, Title: "Buy milk", Done: false},
	{ID: 2, Title: "Write Go code", Done: true},
	{ID: 3, Title: "Read a Go book", Done: false},
	{ID: 4, Title: "Go for a walk", Done: true},
	{ID: 5, Title: "Practice REST API in Go", Done: false},
	{ID: 6, Title: "Call a friend", Done: true},
	{ID: 7, Title: "Learn Docker basics", Done: false},
	{ID: 8, Title: "Clean the desk", Done: false},
	{ID: 9, Title: "Cook dinner", Done: true},
	{ID: 10, Title: "Write a blog post about Go", Done: false},
}

func main() {
	http.HandleFunc("/tasks", getTasks)
	fmt.Println("server started at port : 8081")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		fmt.Println("error in starting server", err)
	}
}

func getTasks(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "No other methods allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("content-type", "application/json")
	jsonData, err := json.Marshal(tasks)

	if err != nil {
		http.Error(w, "error converting into json", http.StatusInternalServerError)
	}

	w.Write(jsonData)
}
