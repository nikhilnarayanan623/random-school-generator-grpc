package models

type School struct {
	Name    string
	Classes []Class
}

type Class struct {
	Name          string
	Students      []Student
	TotalStudents int
}

type Student struct {
	Name           string
	Age            uint
	RollNumber     uint
	Gender         string
	Scores         []Subject
	HaveDisability bool
	Address        Address
}

type Subject struct {
	Name          string
	Score         uint
	Grade         string
	ClassCategory string
	Passed        bool
}

type Address struct {
	HouseNumber int
	City        string
	State       string
}
