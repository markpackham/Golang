package main

import "fmt"

func main() {

	// strings
	var nameOne string = "Hello"
	var nameTwo = "Goodbye"
	var nameThree string
	
	fmt.Println(nameOne, nameTwo, nameThree)

	nameOne = "Jim"
	nameThree = "Bob"

	fmt.Println(nameOne, nameTwo, nameThree)

	nameFour := "Emily"

	fmt.Println(nameOne, nameTwo, nameThree, nameFour)
}