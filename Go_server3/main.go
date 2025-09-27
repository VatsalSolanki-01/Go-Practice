package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	//GET method to get info of all available students
	http.HandleFunc("/students", getStudentsInfo)

	//GET method to get info of specific student based on rollno
	http.HandleFunc("/students", getStudentInfo)
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

}
