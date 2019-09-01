package main

import "fmt"

func main() {
	//Arrays
	// var fruitArr [2]string
	//
	// //Assign values
	// fruitArr[0] = "Apple"
	// fruitArr[1] = "Orange"

	//Declare and assign
	fruitArr := [2]string{"Apple", "Orange"}

	fmt.Println(fruitArr)

	fruitSlice := []string{"Apple", "Orange", "Grape", "Cherry"}

	fmt.Println(fruitSlice)

	//Find the length
	fmt.Println(len(fruitSlice))

	//Get the range
	fmt.Println(fruitSlice[1:3])

}
