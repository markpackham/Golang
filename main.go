package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {

	// Strings package usage (doesn't effect original value)
	greeting := "hello world of meow"
	//  fmt.Println(strings.Contains(greeting, "hello"))
	//  // the original string isn't replaced by ReplaceAll, it uses a new string instead
	//  fmt.Println(strings.ReplaceAll(greeting, "hello","goodbye"))
	//  fmt.Println(strings.ToUpper(greeting))
	//  fmt.Println(strings.Index(greeting, "ll"))
	fmt.Println(strings.Split(greeting, " ")) // [hello world of meow]


	// Sort package usage (DOES effect original value)
	ages := []int{45, 20, 35, 30, 75, 60, 50, 25}

	sort.Ints(ages)
	fmt.Println(ages)

	index := sort.SearchInts(ages, 30)
	fmt.Println(index)

	names := []string{"yoshi", "mario", "peach", "bowser", "luigi"}

	sort.Strings(names)
	fmt.Println(names)

	fmt.Println(sort.SearchStrings(names, "bowser")) // position 0 since it has been sorted

}