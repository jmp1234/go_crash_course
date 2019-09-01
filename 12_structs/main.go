package main

import (
	"fmt"
	"strconv"
)

// Define Person struct
type Person struct {
	// firstName string
	// lastName  string
	// city      string
	// gender    string
	// age       int

	firstName, lastName, city, gender string
	age                               int
}

// Greeting method (value receiver)
func (person Person) greet() string {
	return "Hello, my name is " + person.firstName + " " + person.lastName + " and I am " +
		strconv.Itoa(person.age) + " years old."
}

// hasBirthday method (pointer receiver) - we change something
func (person *Person) hasBirthday() {
	person.age++
}

// getMarried (pointer receiver)
func (person *Person) getMarried(spouseLastName string) {
	if person.gender == "male" {
		return
	} else {
		person.lastName = spouseLastName
	}
}

func main() {
	// Init person using struct
	// person1 := Person{firstName: "Sam", lastName: "Smith",
	// 	city: "Newport Beach", gender: "male", age: 23}

	// Alternative
	person1 := Person{"Sam", "Smith", "Newport Beach", "female", 23}

	fmt.Println(person1)
	fmt.Println(person1.city)
	person1.age++
	fmt.Println(person1)

	person1.hasBirthday()
	person1.getMarried("Williams")

	fmt.Println(person1.greet())

}
