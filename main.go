package main

// a "map" is like a dictionary in python with key value pairs
// in a map it must be all the same type, you can't just chuck anything into it

import(
	"fmt"
)

func main(){
	menu := map[string]float64{
		"soup": 4.99,
		"pie": 7.99,
		"salad": 6.99,
		"toffee": 3.55,
	}

	fmt.Println(menu)
	fmt.Println(menu["pie"])

	// looping maps with key and value (you must use "k" and "v")
	for k, v := range menu{
		fmt.Println(k, "-", v)
	}

	// ints as key type
	phonebook := map[int]string{
		267584967: "mario",
		984759373: "luigi",
		845775485: "peach",
	}

	fmt.Println(phonebook)
	fmt.Println(phonebook[267584967])

	phonebook[984759373] = "wario"
	fmt.Println(phonebook)

	phonebook[845775485] = "toad"
	fmt.Println(phonebook)

}