package main

import (
	"fmt"
	"strings"
)

func getInitials(name string) (string, string) {

	newString := strings.ToUpper(name)
	names := strings.Split(newString, " ")

	var initials []string
	// don't want index but do want value so _, V
	for _, v := range names {
		initials = append(initials, v[:1])
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	}

	// if you only have one name return this
	return initials[0], "_"
}

func main() {
	firstName1, surName1 := getInitials("Squall Leonheart")
	fmt.Println(firstName1, surName1)

	// we can only handle 2 params so "Ardyn Lucis Caelum" would just return "AL" so "C" wouldn't be returned
	fn2, sn2 := getInitials("Ardyn Lucis Caelum")
	fmt.Println(fn2, sn2)

	// Yuna give "Y _"
	fn3, sn3 := getInitials("Yuna")
	fmt.Println(fn3, sn3)
}