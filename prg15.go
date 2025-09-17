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

func (f FullTimeEmployee) createFTE(name string, baseSalary, bonus float64) FullTimeEmployee {
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
