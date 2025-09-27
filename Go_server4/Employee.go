package main

var Employees = []Employee{
	{Name: "John", Id: 1, Salary: 25000},
	{Name: "Bob", Id: 2, Salary: 5000},
	{Name: "Alice", Id: 3, Salary: 55000},
	{Name: "Katrina", Id: 4, Salary: 35000},
	{Name: "Tom", Id: 5, Salary: 25000},
	{Name: "Mike", Id: 6, Salary: 25000},
	{Name: "Dona", Id: 7, Salary: 250000},
	{Name: "Luis", Id: 8, Salary: 25020},
	{Name: "Jessica", Id: 9, Salary: 105000},
	{Name: "Harvey", Id: 10, Salary: 20000},
	{Name: "Leonardo", Id: 11, Salary: 65000},
	{Name: "Thomas", Id: 12, Salary: 250000},
	{Name: "Bob", Id: 13, Salary: 5000},
}

type Employee struct {
	Name   string `json:"name"`
	Id     int    `json:"id"`
	Salary int    `json:"salary"`
}
