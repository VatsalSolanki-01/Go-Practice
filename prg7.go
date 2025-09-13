package main

import "fmt"

var firstMap = make(map[string]int)

func Contacts() {
	for {
		fmt.Println("***************************************************")
		fmt.Println("enter 1 to display contacts")
		fmt.Println("enter 2 to add new contacts")
		fmt.Println("enter 3 to exit the program")
		fmt.Println("***************************************************")
		var input string
		fmt.Scan(&input)
		if input == "3" {
			break
		} else {
			ContactMenu(input)
		}
	}
}

func ContactMenu(input string) {
	switch input {
	case "1":
		displayContactno()
	case "2":
		addContactno()
	}
}
func addContactno() {
	var key string
	var value int
	var n int

	fmt.Println("enter no. of entries to enter :- ")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		fmt.Printf("enter Name:-")
		fmt.Scan(&key)
		fmt.Printf("enter contactno. :-")
		fmt.Scan(&value)

		firstMap[key] = value
	}
}

func displayContactno() {
	for k, v := range firstMap {
		fmt.Println("key :-", k, "value :-", v)
	}
}
