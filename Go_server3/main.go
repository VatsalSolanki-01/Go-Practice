package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	//GET method to get info of all available students
	http.HandleFunc("/students", getStudentsInfo)

	//GET method to get info of specific student based on rollno
	http.HandleFunc("/student", getStudentInfo)
	if err := http.ListenAndServe(":8081", nil); err != nil {
		fmt.Println("error starting server")
	}
}

func getStudentsInfo(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
	}

	w.Header().Set("content-type", "application/json")
	jsonData, err := json.Marshal(studentList)

	if err != nil {
		http.Error(w, "error encoding into json", http.StatusInternalServerError)
	}

	w.Write(jsonData)
}

func getStudentInfo(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
	}

	rollno := r.URL.Query().Get("rollno")

	if rollno == "" {
		http.Error(w, " rollno required ", http.StatusBadRequest)
		return
	}

	roll, err := strconv.Atoi(rollno)
	if err != nil {
		http.Error(w, "invalid rollno ", http.StatusBadRequest)
	}

	for _, v := range studentList {
		if v.Rollno == roll {
			w.Header().Set("Content-type", "application/json")
			json.NewEncoder(w).Encode(v)
			return
		}
	}
	http.Error(w, "student not found", http.StatusNotFound)
}
