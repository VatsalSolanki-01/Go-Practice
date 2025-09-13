package main

import "fmt"

type Address struct {
	cityName string
	pincode  int
}

type Person struct {
	Address
	name string
}

func (p Person) info() {
	fmt.Println("name of the person is :- ", p.name)
	fmt.Println("address and pincode is:-", p.cityName, p.pincode)
}
