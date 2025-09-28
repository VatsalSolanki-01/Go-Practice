package main

type Book struct {
	Title  string `json:"title"`
	Id     int    `json:"id"`
	ISBN   string `json:"isbn"`
	Author string `json:"author"`
}

var Books = []Book{
	{Id: 1, Title: "The Go Programming Language", ISBN: "9780134190440", Author: "Alan A. A. Donovan"},
	{Id: 2, Title: "Clean Code", ISBN: "9780132350884", Author: "Robert C. Martin"},
	{Id: 3, Title: "Design Patterns: Elements of Reusable Object-Oriented Software", ISBN: "9780201633610", Author: "Erich Gamma"},
	{Id: 4, Title: "Introduction to Algorithms", ISBN: "9780262033848", Author: "Thomas H. Cormen"},
	{Id: 5, Title: "The Pragmatic Programmer", ISBN: "9780201616224", Author: "Andrew Hunt"},
	{Id: 6, Title: "Refactoring: Improving the Design of Existing Code", ISBN: "9780201485677", Author: "Martin Fowler"},
	{Id: 7, Title: "Head First Design Patterns", ISBN: "9780596007126", Author: "Eric Freeman"},
	{Id: 8, Title: "Code Complete", ISBN: "9780735619678", Author: "Steve McConnell"},
	{Id: 9, Title: "Working Effectively with Legacy Code", ISBN: "9780131177055", Author: "Michael Feathers"},
	{Id: 10, Title: "Domain-Driven Design", ISBN: "9780321125217", Author: "Eric Evans"},
}
