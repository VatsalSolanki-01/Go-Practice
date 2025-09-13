package main

import "fmt"

type student struct {
	name   string
	rollno int
	marks  int
}

func (s student) Greet() {
	fmt.Println(s.name, "good morning...!!!")
}

func (s *student) changeToFullName(fullname string) {
	s.name = fullname
}

func (s student) info() {
	fmt.Println(s.name)
	fmt.Println(s.rollno)
	fmt.Println(s.marks)
}
