package main

import "fmt"

var eventName string = "GoConference"
var totalTickets int = 100

var bookings []Bookings

type Bookings struct {
	BookedInNameOf string
	noOfTickets    int
}

func (b Bookings) bookTickets(BookedInNameOf string, noOfTickets int) Bookings {
	return Bookings{BookedInNameOf: BookedInNameOf, noOfTickets: noOfTickets}
}

func (b Bookings) Details() {
	fmt.Println("name of the person who booked tickets is :- ", b.BookedInNameOf)
	fmt.Println("no. of tickets are :- ", b.noOfTickets)
}

func GoCOnference() {
	var input string
	for {
		fmt.Println("Hello...!!! Welcome to", eventName)
		fmt.Println("enter 1 to book tickets for GoCOnference:- ")
		fmt.Println("enter 2 to check for remaining seatings :- ")
		fmt.Println("enter 3 to exit menu :-")

		if input == "3" {
			break
		}
	}
	fmt.Scan(&input)
	switch input {
	case "1":
	case "2":
	case "3":
	}
}
