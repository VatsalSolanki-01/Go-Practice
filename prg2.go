package main

import "fmt"

func start() {
	fmt.Println("enter value :- ")
	var input string
	fmt.Scan(&input)

	var value []string
	for _, i := range input {
		value = append(value, string(i))
	}

	if len(value) <= 0 || len(value) > 4 {
		fmt.Println("put valid input")
	}
	if len(value) == 1 {
		fmt.Println(unitsPlace(value[0]))
	}
	if len(value) == 2 {
		fmt.Println(tensPlace(value[0]), unitsPlace(value[1]))
	}
	if len(value) == 3 {
		fmt.Println(hundredsPlace(value[0]), tensPlace(value[1]), unitsPlace(value[2]))
	}
	if len(value) == 4 {
		fmt.Println(thousandsPlace(value[0]), hundredsPlace(value[1]), tensPlace(value[2]), unitsPlace(value[3]))
	}

}

func unitsPlace(input string) string {

	switch input {
	case "1":
		return "i"
	case "2":
		return "ii"
	case "3":
		return "iii"
	case "4":
		return "iV"
	case "5":
		return "V"
	case "6":
		return "Vi"
	case "7":
		return "Vii"
	case "8":
		return "Viii"
	case "9":
		return "iX"
	}
	return ""
}

func tensPlace(input string) string {
	switch input {
	case "1":
		return "X"
	case "2":
		return "XX"
	case "3":
		return "XXX"
	case "4":
		return "XL"
	case "5":
		return "L"
	case "6":
		return "LX"
	case "7":
		return "LXX"
	case "8":
		return "LXXX"
	case "9":
		return "XC"
	}
	return ""
}

func hundredsPlace(input string) string {
	switch input {
	case "1":
		return "C"
	case "2":
		return "CC"
	case "3":
		return "CCC"
	case "4":
		return "CD"
	case "5":
		return "D"
	case "6":
		return "DC"
	case "7":
		return "DCC"
	case "8":
		return "DCCC"
	case "9":
		return "CM"
	}
	return ""
}

func thousandsPlace(input string) string {
	switch input {
	case "1":
		return "M"
	case "2":
		return "MM"
	case "3":
		return "MMM"
	}
	return ""
}
