package main

import(
	"fmt"
)

var score = 99.5

func main() {
	sayHello("mario")

	for _, v := range points{
		fmt.Println(v)
	}

	showScore()
}


// To run this file you must also run ones it is using
// go run main.go greetings.go 