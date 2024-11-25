package main

import "fmt"

type FullName struct {
	// TODO: add fields
	firstName string
	lastName  string
}

// TODO: declare a structure for birth date
type birthdate struct {
	dayOfBirth   int
	monthOfBirth int
	yearOfBirth  int
}

type Profile struct {
	// TODO: embed full name and birth date information
	NumberOfSiblings byte
	ZodiacSign       rune
}

func main() {
	var me = Profile{
		// TODO: set name and birth date information
		NumberOfSiblings: 2,        // TODO: adjust
		ZodiacSign:       '\u264C', // TODO: adjust
	}
	fmt.Println(me)

	fmt.Println("Siblings Before:", me.NumberOfSiblings)
	// TODO: imagine, you get a little brother or sister
	me.NumberOfSiblings++
	fmt.Println("Siblings After:", me.NumberOfSiblings)
}
