package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	//Get method Greetings page
	http.HandleFunc("/", Greetings)
	// Get method Geting all books
	http.HandleFunc("/getbooks", getBooks)

	//Get method getting book by id
	http.HandleFunc("/getbookid", getBookById)

	//Post method adding book
	http.HandleFunc("/addbook", addBook)

	fmt.Println("server running on port : 8081")
	if err := http.ListenAndServe("localhost:8081", nil); err != nil {
		fmt.Println("error starting server", err)
	}
}

func Greetings(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Welcome to Library</h1>")
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-type", "application/json")
	jsonData, err := json.Marshal(Books)
	if err != nil {
		http.Error(w, "issue generating json", http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}

func getBookById(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	book_id := r.URL.Query().Get("id")

	if book_id == "" {
		http.Error(w, "id needed to fetch book", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(book_id)
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	for _, v := range Books {
		if v.Id == id {
			w.Header().Set("Content-type", "application/json")
			json.NewEncoder(w).Encode(v)
			return
		}
	}

	http.Error(w, "book not found", http.StatusNotFound)
}

func addBook(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not aloowed", http.StatusMethodNotAllowed)
		return
	}

	var newBook Book
	err := json.NewDecoder(r.Body).Decode(&newBook)

	if err != nil {
		http.Error(w, "proper data needed", http.StatusBadRequest)
		return
	}

	if newBook.Author == "" || newBook.ISBN == "" || newBook.Id == 0 || newBook.Title == "" {
		http.Error(w, "proper data needed", http.StatusBadRequest)
		return
	}

	Books = append(Books, newBook)

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(newBook)
}
