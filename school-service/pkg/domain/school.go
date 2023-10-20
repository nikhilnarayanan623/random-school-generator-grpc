package domain

type State struct {
	Name   string
	Cities []string
}


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

type Address struct {
	HouseNumber int
	City        string
	State       string
}
