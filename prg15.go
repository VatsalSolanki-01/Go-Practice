package main

import "fmt"

var Employeelist []Employee

func addEmployee() {
	var input string
	fmt.Println("enter a to add Full-Time-Employee and b to add Part-Time-Employee :- ")
	fmt.Scan(&input)

	switch input {
	case "a":
		var name string
		var baseSalary, bonus float64
		fmt.Println("enter name of Employee :- ")
		fmt.Scan(&name)
		fmt.Println("enter base Salary :- ")
		fmt.Scan(&baseSalary)
		fmt.Println("enter bonsu :- ")
		fmt.Scan(&bonus)
		var f FullTimeEmployee
		Employeelist = append(Employeelist, f.createFTE(name, baseSalary, bonus))
	case "b":
		var name string
		var HourlyRate, HoursWorked float64
		fmt.Println("enter name of the employee :- ")
		fmt.Scan(&name)
		fmt.Println("enter hourly rate of employee :- ")
		fmt.Scan(&HourlyRate)
		fmt.Println("enter hours worked :- ")
		fmt.Scan(&HoursWorked)

		var p PartTimeEmployee
		Employeelist = append(Employeelist, p.createPTE(name, HourlyRate, HoursWorked))
	}
}

func displayEmployees() {
	for _, v := range Employeelist {
		v.display()
	}
}

type Employee interface {
	salary() float64
	display()
}

func salary(e Employee) float64 {
	return e.salary()
}

func display(e Employee) {
	e.display()
}

type FullTimeEmployee struct {
	name       string
	baseSalary float64
	bonus      float64
}

func (f FullTimeEmployee) salary() float64 {
	return f.baseSalary + f.bonus
}

func (f FullTimeEmployee) display() {
	fmt.Println("name of the employee :- ", f.name)
	fmt.Println("salary of the employee :- ", f.salary())
}

func (f FullTimeEmployee) createFTE(name string, baseSalary float64, bonus float64) FullTimeEmployee {
	return FullTimeEmployee{name: name, baseSalary: baseSalary, bonus: bonus}
}

type PartTimeEmployee struct {
	name        string
	HourlyRate  float64
	HoursWorked float64
}

func (p PartTimeEmployee) salary() float64 {
	return p.HourlyRate * p.HoursWorked
}

func (p PartTimeEmployee) display() {
	fmt.Println("name of the employee :- ", p.name)
	fmt.Println("salary of the employee :- ", p.salary())
}

func (p PartTimeEmployee) createPTE(name string, HourlyRate, HoursWorked float64) PartTimeEmployee {
	return PartTimeEmployee{name: name, HourlyRate: HourlyRate, HoursWorked: HoursWorked}
}
