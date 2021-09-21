package main

import "fmt"

// Go uses Struct for custom types, rather than Classes
func main(){

	myBill := newBill("mario bill")

	fmt.Println(myBill)
}


// go run main.go bill.go