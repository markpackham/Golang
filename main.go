package main

import "fmt"

func updateName(x string){
	x = "wedge"
}

func main(){

	name := "tifa"

	updateName(name)

	fmt.Println("memory address of name is: ", &name)

	m := &name

	fmt.Println(name)
}




/*
Memory addresses
|--name---|----m----|
|  0x001  |  0x002  |
|---------|---------|
| "tifa"  | p0x001  |
|---------|---------|
*/