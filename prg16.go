package main

import "fmt"

var eventName string = "GoConference"
var totalTickets int = 100
var bookings []Bookings

type Bookings struct {
	BookedInNameOf string
	noOfTickets    int
	email          string
}

func (b Bookings) bookTickets(BookedInNameOf string, noOfTickets int, email string) Bookings {
	return Bookings{BookedInNameOf: BookedInNameOf, noOfTickets: noOfTickets, email: email}
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
		fmt.Scan(&input)

		if input == "3" {
			break
		}

		switch input {
		case "1":
			ticketBooking()
		case "2":
		}
	}

}

func ticketBooking() {
	var name string
	var email string
	var noOfTickets int
	fmt.Println("Please enter your name :- ")
	fmt.Scan(&name)
	fmt.Println("please enter your email :- ")
	fmt.Scan(&email)
	fmt.Println("enter no. of tickets to book :- ")
	fmt.Scan(&noOfTickets)

	var booking Bookings
	if totalTickets >= noOfTickets {
		bookings = append(bookings, booking.bookTickets(name, noOfTickets, email))
		totalTickets -= noOfTickets
	} else {
		fmt.Println("sorry seats are full will let you know next time conference is held ")
	}
}
