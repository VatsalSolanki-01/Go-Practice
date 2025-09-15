package main

type shape interface {
	area() float64
	perimeter() float64
}

func area(sh shape) float64 {
	return sh.area()
}

func perimeter(sh shape) float64 {
	return sh.perimeter()
}

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

type Square struct {
	side float64
}

func (s Square) area() float64 {
	return s.side * s.side
}

func (s Square) perimeter() float64 {
	return 4 * s.side
}
