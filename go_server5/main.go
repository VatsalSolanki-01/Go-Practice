package main

import (
	"fmt"
	"net/http"
)

// var Counter int = 0
var counter = make(chan int, 50) //channel having 100 buffers if the race condition arrives
var countV int = 0

func count() {
	for range counter {
		countV += 1
	}
}
func main() {
	go count()
	http.HandleFunc("/hello", helloWorld)
	http.HandleFunc("/getcount", getcount)
	fmt.Println("server started at 8081")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		fmt.Println("server not started", err)
	}
}

func getcount(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprintf(w, "<h1>%d</h1>", countV)
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	counter <- 1

	fmt.Fprintf(w, "<h1> Hello World !!!</h1>")

}
