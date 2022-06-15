package main

import "fmt"

type Person struct {
	name    string
	age     int
	friends []Person
}

func main() {

	firstPerson := Person{
		name:    "Adi",
		age:     30,
		friends: nil,
	}

	secondPerson := Person{
		name:    "Mircea",
		age:     10,
		friends: []Person{firstPerson},
	}
	fmt.Println(firstPerson)
	fmt.Println(secondPerson)
}
