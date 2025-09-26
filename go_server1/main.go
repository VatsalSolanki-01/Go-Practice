package main

import (
	"fmt"
	"net/http"
)

func Greetings(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>good evening</h1>")
}

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1> hello world !!!</h1>")
}

func main() {
	http.HandleFunc("/", Greetings)

	http.HandleFunc("/home", Home)

	fmt.Println("server running on port : 8081")
	if err := http.ListenAndServe("localhost:8081", nil); err != nil {
		fmt.Println("error starting server", err)
	}
}
