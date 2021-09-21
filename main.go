package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func getInput(prompt string, r *bufio.Reader) (string, error){
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func createBill() bill{
	reader := bufio.NewReader(os.Stdin)

	// fmt.Print("Create a new bill name: ")
	// name, _ := reader.ReadString('\n')
	// name = strings.TrimSpace(name)

	name, _ := getInput("Create a new bill name: ", reader)

	b := newBill(name)
	fmt.Println("Created the bill - ", b.name)

	return b
}

func promptOptions(b bill){
	reader := bufio.NewReader(os.Stdin)

	opt, _ := getInput("Choose option (a - add item, s - save bill, t - add tip): ", reader)

	switch opt{
	case "a":
		name, _ := getInput("Item name: ", reader)
		price, _ := getInput("Item price: ", reader)

		p, err := strconv.ParseFloat(price, 64)
		if err != nil{
			fmt.Println("Price must be a number")
			promptOptions(b)
		}
		b.addItem(name, p)

		fmt.Println("item added - ", name, p)
		promptOptions(b)
	case "t":
		tip, _ := getInput("Enter tip amount ($): ", reader)

		t, err := strconv.ParseFloat(tip, 64)
		if err != nil{
			fmt.Println("Tip must be a number")
			promptOptions(b)
		}
		b.updateTip(t)

		fmt.Println("Tip updated - ", t)
		promptOptions(b)
	case "s":
		fmt.Println("You chose to save the bill", b)
	default:
		fmt.Println("That was not a valid option...")
		promptOptions(b)
	}
}

func main(){
	mybill := createBill()
	promptOptions(mybill)
}


// go run main.go bill.go