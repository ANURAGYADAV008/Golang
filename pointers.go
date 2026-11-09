package main

import "fmt"

type student struct {
	firstName string
	lastName  string
	id        int
}

func pointerkaro() {
	s := student{
		firstName: "Anurag",
		lastName:  "Yadav",
		id:        123,
	}
	fmt.Println(s)
	updatestudent(&s, "Jay")
	fmt.Print(s)

}

func increaseval(x *int) {
	*x++
}
func updatestudent(s *student, lastName string) {
	s.lastName = lastName
}
