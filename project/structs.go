package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')
	return strings.TrimSpace(input), err
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	// fmt.Print("Create a new bill name: ")
	// name, _ := reader.ReadString('\n')
	// name = strings.TrimSpace(name)

	name, _ := getInput("Create a new bill name: ", reader)

	b := newBill(name)
	fmt.Println("Created the bill - ", b.name)

	return b
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInput("Choose option (a - add an Item, s - save bill, t - add a tip) ", reader)

	switch opt {
	case "a":
		name, _ := getInput("Item name: ", reader)
		price, _ := getInput("Item price: ", reader)
		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("the price must be a number")
			promptOptions(b)
		}
		b.addItem(name, p)
		fmt.Println("Item added", name, p)
		promptOptions(b)
	case "t":
		tip, _ := getInput("Item price: ", reader)
		t, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("the price must be a number")
			promptOptions(b)
		}
		b.updateTip(t)
		fmt.Println("tip added", t)
		promptOptions(b)
	case "s":
		b.save()
		fmt.Println("You saved the file", b.name+".txt")
	default:
		fmt.Println("that was not a valid option...")
		promptOptions(b)
	}
}

func main() {
	myBill := createBill()
	promptOptions(myBill)
	// myBill.addItem("coffee Puding", 20.99)
	// myBill.updateTip(10)

	fmt.Println(myBill)
}
