package main

import (
	"fmt" 
	"os"
)

// Go uses Struct for custom types, rather than Classes

type bill struct {
	name string
	items map[string]float64
	tip float64
}

// make new bills
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}
	return b
}


// receiver functions are like methods for classes
// this function will only work with bill objects
func (b *bill) format() string{
	fs := "Bill breakdown: \n"
	var total float64 = 0

	// list items
	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v £%v \n", k+":", v)
		total += v
	}

	// add tip
	fs += fmt.Sprintf("%-25v £%v\n", "tip:", b.tip)

	// total
	fs += fmt.Sprintf("%-25v £%0.2f", "total:", total+b.tip)

	return fs
}

// update tip (need to use pointer in argument to make a change)
func (b *bill) updateTip(tip float64){
	// you don't need to bother dereferencing pointers with structs
	(b).tip = tip
}

// add an item to the bill
func (b *bill) addItem(name string, price float64){
	b.items[name] = price
}

// save bill, 0644 are the file permissions
func (b *bill) save(){
	data := []byte(b.format())

	err := os.WriteFile("bills/"+b.name+".txt", data, 0644)

	if err != nil{
		// panic kills the program
		panic(err)
	}
	fmt.Println("Bill was saved to file")
}