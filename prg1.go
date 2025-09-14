package main

import "fmt"

func main() {
	arr := make([]int, 10)
	target := 0
	fmt.Println("enter elements :- ")

	for i := 0; i <= 5; i++ {
		fmt.Scan(&arr[i])
	}
	fmt.Println("enter target value:-")
	fmt.Scan(&target)

	fmt.Println(twoSum(arr, target))
}
