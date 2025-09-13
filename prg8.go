package main

import "fmt"

func practiceSlice() {
	var arr []int

	n := 5
	var input int
	for i := 0; i < n; i++ {
		fmt.Scan(&input)
		arr = append(arr, input)
	}

	for _, value := range arr {
		fmt.Println("value :- ", value)
	}
}
