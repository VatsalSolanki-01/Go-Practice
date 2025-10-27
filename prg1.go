package main

import "fmt"

func main() {
	messagePipe := make(chan string)

	fmt.Println(<-messagePipe)
	fmt.Println(<-messagePipe)
	messagePipe <- "message from main goroutine"
	messagePipe <- "2nd message from main go routine"

}
