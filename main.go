package main

import "fmt"

func main() {

	age := 10
	name := "Bob"

	// // Print
	// fmt.Print("Hello ")
	// fmt.Print("world \n")
	// fmt.Print("world \n")

	// // PrintLn
	// fmt.Println("Hello World")
	// fmt.Println("Hello World")
	// fmt.Println("My age is", age, " my name is ", name)

	// Printf (Formatted string) %_ = format specifier
	fmt.Printf("my age is %v and my name is %v \n", age, name)
	fmt.Printf("my age is %d and my name is %q \n", age, name)
	fmt.Printf("age is of type %T", age)
	fmt.Printf("you scored %0.2f points! \n", 222.3311)


	// Sprintf (save formatted strings in variables)
	var str = fmt.Sprintf("my age is %v and my name is %v \n", age, name)
	fmt.Println("The saved String is -- ", str)

}