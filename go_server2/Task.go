package main

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

var Tasks = []Task{
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
