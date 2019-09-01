package main

import "fmt"

func main() {
	// Define map
	// emails := make(map[string]string)

	// Assign key-values
	// emails["Bob"] = "bob@gmail.com"
	// emails["Tim"] = "tim@gmail.com"
	// emails["Mike"] = "mike@gmail.com"

	//Declare map and add key-values
	emails := map[string]string{
		"Bob":  "bob@gmail.com",
		"Tim":  "tim@gmail.com",
		"Mike": "mike@gmail.com",
	}

	emails["John"] = "john@gmail.com"

	fmt.Println(emails)
	fmt.Println(emails["Bob"])
	fmt.Println(len(emails))

	// Delete from map
	delete(emails, "Bob")
	fmt.Println(emails)
}
