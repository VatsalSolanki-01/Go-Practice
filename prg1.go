package main

import "fmt"

func main() {
	r := Rectangle{length: 10, width: 5}
	s := Square{side: 7}

	fmt.Println("Rectangle Area:", area(r))
	fmt.Println("Rectangle Perimeter:", perimeter(r))

	fmt.Println("Square Area:", area(s))
	fmt.Println("Square Perimeter:", perimeter(s))
}
