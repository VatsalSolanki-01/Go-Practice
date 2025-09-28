package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	//Reading
	http.HandleFunc("/tasks", getTasks)

	//Creading
	http.HandleFunc("/addtask", addTask)

	//Updating
	http.HandleFunc("/updatetask", updateTask)

	//Deleting
	http.HandleFunc("/deletetask", deleteTask)

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
	jsonData, err := json.Marshal(Tasks)

	if err != nil {
		http.Error(w, "error converting into json", http.StatusInternalServerError)
	}

	w.Write(jsonData)
}

func addTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var newTask Task
	json.NewDecoder(r.Body).Decode(&newTask)

	if newTask.Title == "" || newTask.ID == 0 {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	Tasks = append(Tasks, newTask)

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(newTask)
}

func updateTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	task_id := r.URL.Query().Get("id")

	if task_id == "" {
		http.Error(w, "need id to update", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(task_id)

	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	var newTask Task
	json.NewDecoder(r.Body).Decode(&newTask)

	if newTask.Title == "" || newTask.ID == 0 {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	flag := false

	for i, v := range Tasks {
		if v.ID == id {
			flag = true
			Tasks[i].Title = newTask.Title
			Tasks[i].Done = newTask.Done

			w.Header().Set("Content-type", "application/json")
			json.NewEncoder(w).Encode(Tasks[i])
			return
		}
	}
	if flag != true {
		http.Error(w, "task not found", http.StatusNotFound)
		return
	}
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	task_id := r.URL.Query().Get("id")

	if task_id == "" {
		http.Error(w, "need id to delete task", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(task_id)

	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	flag := false
	for i, v := range Tasks {
		if v.ID == id {
			flag = true
			deletedTask := Tasks[i]
			Tasks = append(Tasks[:i], Tasks[i+1:]...)
			w.Header().Set("Content-type", "application/json")
			json.NewEncoder(w).Encode(deletedTask)
			return
		}
	}
	if flag != true {
		http.Error(w, "task not found", http.StatusNotFound)
	}
}
