package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("function hello :- hello-1")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("function hello :- hello-2")
}

func hi() {
	fmt.Println("function hi :- hi-1")
}
