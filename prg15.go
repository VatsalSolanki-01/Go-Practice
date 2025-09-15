package main

import "fmt"

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

type FullTime struct {
	name       string
	baseSalary float64
	bonus      float64
}

func (f FullTime) salary() float64 {
	return f.baseSalary + f.bonus
}

func (f FullTime) display() {
	fmt.Println("name of the employee :- ", f.name)
	fmt.Println("salary of the employee :- ", f.salary())
}

type PartTime struct {
	name        string
	HourlyRate  float64
	HoursWorked float64
}

func (p PartTime) salary() float64 {
	return p.HourlyRate * p.HoursWorked
}

func (p PartTime) display() {
	fmt.Println("name of the employee :- ", p.name)
	fmt.Println("salary of the employee :- ", p.salary())
}
