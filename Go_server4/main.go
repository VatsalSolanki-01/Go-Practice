package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	// landing page
	http.HandleFunc("/", Greetings)

	//Get all employees
	http.HandleFunc("/employees", getAllEmployees)

	//Get Employee by id
	http.HandleFunc("/employeeid", getEmployeeById)

	//Get Employee by name
	http.HandleFunc("/employeename", getEmployeeByName)

	//Post Method to add employee
	http.HandleFunc("/addemployees", addEmployee)

	fmt.Println("server started at port :9191")
	if err := http.ListenAndServe(":9191", nil); err != nil {
		fmt.Println("error starting server", err)
	}
}

func Greetings(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
	}

	fmt.Fprintf(w, "<h1>Greetings welcome to Employee Department </h1>")
}

func getAllEmployees(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	w.Header().Set("Content-type", "application/json")
	jsonData, err := json.Marshal(Employees)

	if err != nil {
		http.Error(w, "cannot convert into json", http.StatusInternalServerError)
	}

	w.Write(jsonData)
}

func getEmployeeById(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	e_id := r.URL.Query().Get("ID")

	if e_id == "" {
		http.Error(w, "id needed", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(e_id)

	if err != nil {
		http.Error(w, "proper id is needed", http.StatusBadRequest)
		return
	}

	for _, v := range Employees {
		if v.Id == id {
			w.Header().Set("Content-type", "application/json")
			json.NewEncoder(w).Encode(v)
			return
		}
	}
	http.Error(w, "Employee not found haha", http.StatusNotFound)
}

func getEmployeeByName(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	E_name := r.URL.Query().Get("name")

	if E_name == "" {
		http.Error(w, "name needed", http.StatusBadRequest)
		return
	}

	for _, v := range Employees {
		if v.Name == E_name {
			w.Header().Set("Content-type", "application/json")
			json.NewEncoder(w).Encode(v)
			return
		}
	}
	http.Error(w, "Employee not found hahaha", http.StatusNotFound)
}

func addEmployee(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var newEmp Employee
	err := json.NewDecoder(r.Body).Decode(&newEmp)

	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	if newEmp.Id == 0 || newEmp.Salary == 0 || newEmp.Name == "" {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	Employees = append(Employees, newEmp)

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newEmp)
}
