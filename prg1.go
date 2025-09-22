package main

import "time"

func main() {
	go hello()
	go hi()

	time.Sleep(1001 * time.Millisecond)
}
