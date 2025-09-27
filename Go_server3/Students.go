package main

var studentList = []Student{
	{Rollno: 1, Name: "abcd"},
	{Rollno: 2, Name: "bcde"},
	{Rollno: 3, Name: "cdef"},
	{Rollno: 4, Name: "defg"},
	{Rollno: 5, Name: "efgh"},
	{Rollno: 6, Name: "fghi"},
	{Rollno: 7, Name: "ghij"},
	{Rollno: 8, Name: "hijk"},
	{Rollno: 9, Name: "ijkl"},
	{Rollno: 10, Name: "jklm"},
}

type Student struct {
	Rollno int    `json:"rollno"`
	Name   string `json:"name"`
}
