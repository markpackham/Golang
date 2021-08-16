package main

import (
	"fmt"
)

func main() {

//  age := 450

//  if age < 30{
// 	 fmt.Println("Age less than 30")
//  } else if age < 40{
// 	fmt.Println("Age less than 40")
//  } else{
// 	fmt.Println("Age more than 39") 
//  }

 names := []string{"mario", "luigi", "yoshi", "peach", "bowser"}

 for index, val := range names {
	if index == 1 {
		fmt.Println("continuing at pos", index)
		continue
	}
	if index > 2 {
		fmt.Println("breaking at pos", index)
		break
	}
	fmt.Printf("the value at pos %v is %v \n", index, val)

// Output
// the value at pos 0 is mario 
// continuing at pos 1
// the value at pos 2 is yoshi
// breaking at pos 3


 }



}