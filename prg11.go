package main

import "fmt"

// Problem Definition:
// You are given a list of integers that represent the daily gold you collected in a game for 7 days.
// Store them in an array.
// Convert that array into a slice.
// Find and print:
// The total gold collected in the week.
// The maximum gold collected in a single day.
// A slice containing only the days where gold collected was more than 50.

func prg11Start() {
	var arr [7]int
	var input int
	fmt.Println("enter gold collection of 7 different days :- ")
	var maximumGold int = arr[0]
	var totalGold int = 0
	var goldMoreThan50Days []int

	for i := 0; i <= 6; i++ {
		fmt.Scan(&input)
		arr[i] = input
	}

	slice := arr[0:7]

	fmt.Println("the gold collection day wise is :- ")
	fmt.Println("*****************************************")
	for index, value := range slice {
		fmt.Println("Day :- ", index, "gold amount :- ", value)
	}

	for i := 0; i <= 6; i++ {
		if arr[i] > maximumGold {
			maximumGold = arr[i]
		}
	}

	for i := 0; i <= 6; i++ {
		totalGold += arr[i]
	}

	for i := 0; i <= 6; i++ {
		if arr[i] >= 50 {
			goldMoreThan50Days = append(goldMoreThan50Days, i)
		}
	}
	fmt.Println("*****************************************")
	fmt.Println("total amount of gold collected :-", totalGold)
	fmt.Println("maximum gold collected is :- ", maximumGold)

	for _, value := range goldMoreThan50Days {
		fmt.Println("the days in which gold collection was more than 50 are :- ", value)
	}
	fmt.Println("*****************************************")

}
