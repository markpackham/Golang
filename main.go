package main

import(
	"fmt"
)

func updateName(x string) string{
	x = "wedge"
	return x
}

// effects variables who you can target in memory to change
func updateMenu(y map[string]float64) {
	y["coffee"] = 2.99
}

func main(){
	// group A types -> strings, ints, bools, floats, arrays, structs
	// non-pointer wrapper values
	name := "tifa"

	updateName(name)

	fmt.Println(name)

	// group B types -> slices, maps, functions
	// pointer wrapper values
	menu := map[string]float64{
		"pie":       5.95,
		"ice cream": 3.99,
	}

	updateMenu(menu)
	fmt.Println(menu)

}