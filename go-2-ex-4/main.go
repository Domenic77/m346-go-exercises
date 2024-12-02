package main

import (
	"fmt"
)

func main() {
	// TODO: declare a type for Student (with first and last name)
	type Student struct {
		FirstName string
		LastName  string
	}
	// TODO: declare a type for Class (consisting of multiple students)
	type Class struct {
		Name     string
		Students []Student
	}

	student1 := Student{FirstName: "Riccardo", LastName: "Greco"}
	student2 := Student{FirstName: "Steven", LastName: "Mattman"}

	// TODO: declare a map of modules being attended by multiple classes
	modules := map[int]string{
		104: "Modul 104 ist genial",
		117: "Modul 117 ist hammer",
		346: "Modul 346 ist mega",
	}
	// TODO: output everything using fmt.Println()
	fmt.Println("Modul 104:", modules[104])
	fmt.Println("Modul 117:", modules[117])
	fmt.Println("Modul 346:", modules[346])

	fmt.Println("Student 1:", student1.FirstName, student1.LastName)
	fmt.Println("Student 2:", student2.FirstName, student2.LastName)
}
