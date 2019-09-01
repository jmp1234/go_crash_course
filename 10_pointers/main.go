package main

import "fmt"

func main() {
	a := 5
	b := &a

	//print out 5 and memory address of a
	fmt.Println(a, b)

	//check the type
	fmt.Printf("%T\n", b)

	//Use * to read val from address
	fmt.Println(*b)

	// Change val with pointer
	*b = 10
	fmt.Println(a)
}
