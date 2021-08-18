package main

import (
	"fmt"
	"math"
)

func sayGreeting(name string) {
	fmt.Printf("Good morning %v \n", name)
}

func sayBye(name string) {
	fmt.Printf("Goodbye %v \n", name)
}

// we can pass functions as arguments
func cycleNames(name []string, fun func(string)) {
	// we don't care about the index so user _
	for _, element := range name {
		fun(element)
	}
}


func circleArea(radius float64) float64 {
	// had to import math package to use Pi
	return math.Pi * radius * radius
}

func main() {
	sayGreeting("mario")
	sayGreeting("luigi")
	sayBye("mario")

	cycleNames([]string{"squall", "yuna", "ardyn"}, sayGreeting)
	cycleNames([]string{"squall", "yuna", "ardyn"}, sayBye)
	
	a1 := circleArea(10.5)
	a2 := circleArea(15)

	fmt.Println(a1, a2)
	fmt.Printf("circle 1 area is %0.3f & circle 2 area is %0.3f \n", a1, a2)
}