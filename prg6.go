package main

import "fmt"

// Requirements :-

// Book struct
// Fields: Title string, Author string, ISBN string
// Method: Details() → prints "Title: <Title>, Author: <Author>, ISBN: <ISBN>"
// Member struct
// Fields: Name string, MemberID int
// Method: Introduce() → prints "Member: <Name>, ID: <MemberID>"
// BorrowRecord struct (outer struct)
// Fields: RecordID int, BorrowDate string
// Embed both Book and Member
// Method: ShowRecord() → prints record ID and borrow date, then calls:
// Book.Details()
// Member.Introduce()

type Books struct {
	title  string
	author string
	isbn   string
}

func (b Books) Details() {
	fmt.Println("title of the book:-", b.title)
	fmt.Println("author of the book:-", b.author)
	fmt.Println("isbn number of the book:-", b.isbn)
}

type Member struct {
	name     string
	memberId int
}

func (m Member) Introduce() {
	fmt.Println("name of the memeber is :- ", m.name)
	fmt.Println("member id of the memeber is :- ", m.memberId)
}

type BorrowRecord struct {
	RecordID   int
	BorrowDate string
	Books
	Member
}

func (br BorrowRecord) ShowRecord() {
	fmt.Println("record id is :- ", br.RecordID)
	fmt.Println("borrowing date is :- ", br.BorrowDate)
	br.Details()
	br.Introduce()
}
