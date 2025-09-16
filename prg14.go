package main

import (
	"fmt"
	"math"
)

// slice of type shape
var shapes []shape

func ShapePlay() {
	var input string
	for {
		fmt.Println("enter 1 to create shape ")
		fmt.Println("enter 2 to display details of all shapes ")
		fmt.Println("enter 3 to exit ")
		fmt.Scan(&input)
		if input == "3" {
			break
		} else {
			ShapePlayOptions(input)
		}
	}
}

func ShapePlayOptions(input string) {
	switch input {
	case "1":
		createShapes()
	case "2":
		displayShapes()
	}
}

// to create shapes and append in the list of shapes slice
func createShapes() {
	var input string
	fmt.Println("enter a to create Rectangle b to create Square and c to create Triangle")
	fmt.Scan(&input)
	switch input {
	case "a":
		var length float64
		var width float64
		fmt.Println("enter length:-")
		fmt.Scan(&length)
		fmt.Println("enter width:-")
		fmt.Scan(&width)
		var r Rectangle
		shapes = append(shapes, r.createRectangle(length, width))

	case "b":
		var side float64
		fmt.Println("enter side:-")
		fmt.Scan(&side)
		var s Square
		shapes = append(shapes, s.createSquare(side))

	case "c":
		var a, b, c float64
		fmt.Print("enter a :- ")
		fmt.Scan(&a)
		fmt.Print("enter b :- ")
		fmt.Scan(&b)
		fmt.Print("enter c :- ")
		fmt.Scan(&c)
		var t Triangle
		shapes = append(shapes, t.createTriangle(a, b, c))
	}
}

// display shapes' details
func displayShapes() {
	for _, ob := range shapes {
		ob.details()
	}
}

// interface
type shape interface {
	area() float64
	perimeter() float64
	details()
}

func area(sh shape) float64 {
	return sh.area()
}

func perimeter(sh shape) float64 {
	return sh.perimeter()
}

func details(sh shape) {
	sh.details()
}

// Rectangle structure
type Rectangle struct {
	length float64
	width  float64
}

func (r Rectangle) area() float64 {
	return r.length * r.width
}

func (r Rectangle) perimeter() float64 {
	return 2 * (r.length + r.width)
}

func (r Rectangle) details() {
	fmt.Println("Shape type = Rectangle")
	fmt.Println("length :- ", r.length, "width :- ", r.width)
}

// constructor like method which returns the Rectangle
func (r Rectangle) createRectangle(length_i, width_i float64) Rectangle {
	return Rectangle{length: length_i, width: width_i}
}

// Square structure
type Square struct {
	side float64
}

func (s Square) area() float64 {
	return s.side * s.side
}

func (s Square) perimeter() float64 {
	return 4 * s.side
}

// constructor like method which returns the Square object
func (s Square) createSquare(side float64) Square {
	return Square{side: side}
}

func (s Square) details() {
	fmt.Println("Shape type = Square")
	fmt.Println("side's length :- ", s.side)
}

type Triangle struct {
	a float64
	b float64
	c float64
}

func (t Triangle) area() float64 {
	s := (t.a + t.b + t.c) / 2
	return math.Sqrt(s * (s - t.a) * (s - t.b) * (s - t.c))
}

func (t Triangle) perimeter() float64 {
	return t.a + t.b + t.c
}

func (t Triangle) details() {
	fmt.Println("Shape is Triangle")
	fmt.Println("sides are :-", t.a, t.b, t.c)
}

func (t Triangle) createTriangle(a, b, c float64) Triangle {
	return Triangle{a: a, b: b, c: c}
}
