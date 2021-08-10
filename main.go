package main

import "fmt"

func main() {

	// var ages [3]int = [3]int{20, 25, 30}
	var ages = [3]int{20, 25, 30}

	names := [4]string{"yoshi", "mario", "peach", "bowser"}
	names[1] = "luigi"

	fmt.Println(ages, len(ages))
	fmt.Println(names, len(names))

	// slices (use arrays under the hood but length can change)
	var scores = []int{100, 50, 60}
	scores[2] = 25
	scores = append(scores, 85)

	fmt.Println(scores, len(scores))
	// gives, [100 50 25 85] 4


	// slice ranges (inclusive of value on left but not right)
	rangeOne := names[1:4] // doesn't include pos 4 element
	rangeTwo := names[2:]  //includes the last element
	rangeThree := names[:3] // go from the start and get up to but not including position 3

	fmt.Println(rangeOne, rangeTwo, rangeThree)
	// [luigi peach bowser] [peach bowser] [yoshi luigi peach]
	fmt.Printf("the type of rangeOne is %T \n", rangeOne)

	rangeOne = append(rangeOne, "koopa")
	fmt.Println(rangeOne)


}