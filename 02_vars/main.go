package main

import "fmt"

func main() {
	//MAIN TYPES
	//string
	//bool
	//int
	// int int8 int16 int32 int64
	//uint uint8 uint16 uint32 uint64 uintptr
	//byte - alias for uint8
	//rune - alias for int32
	//float32 float64
	//complex64 complex128

	//Using var
	// var name = "Johnny"
	var age int32 = 23
	const isCool = true
	var size float32 = 2.3

	//Shorthand - can only run within a function
	// name := "Johnny"
	// email := "johnny@gmail.com"

	name, email := "Johnny", "johnny@gmail.com"

	fmt.Println(name, age, isCool, email)
	//print the type
	fmt.Printf("%T\n", size)
}
