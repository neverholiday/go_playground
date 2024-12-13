package main

import (
	"fmt"

	"github.com/google/uuid"
)

type Students struct {
	Id   uuid.UUID
	Name string
}

func ProcessStudent(s Students) {
	fmt.Printf("addr: %p\n", &s)
	fmt.Println("Id : ", s.Id)
	fmt.Println("Name : ", s.Name)
	s.Id = uuid.New()
	s.Name = "Newbie"
}

func foo(s []Students) {
	fmt.Printf("addr: %p\n", &s)
}

func main() {

	fmt.Println("Hello world")
	var students []Students
	students = append(students, Students{
		Name: "Nas",
		Id:   uuid.New(),
	})

	foo(students)
	fmt.Printf("addr: %p\n", &students)

}
