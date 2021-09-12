package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(prompt string) (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	input, err := reader.ReadString('\n')

	return strings.TrimSpace(input), err
}

func promptOptions(bill Bill) {
	input, _ := getInput("Choose option (a - add item, s - save bill, t - add tip): ")
	fmt.Println(input)

	switch input {
	case "a":
		name, _ := getInput("Item name: ")
		price, _ := getInput("Item price: ")

		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("The price must be a number...")
			promptOptions(bill)
		}
		bill.addItem(name, p)

		fmt.Println("item added -", name, price)
		promptOptions(bill)
	case "t":
		tip, _ := getInput("Enter tip amount ($): ")

		t, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("The tip must be a number...")
			promptOptions(bill)
		}
		bill.updateTip(t)

		fmt.Println("tip has been updated to", tip)
		promptOptions(bill)
	case "s":
		bill.save()
		fmt.Println("bill has been saved as", bill.name)
	default:
		fmt.Println("That was not a valid option...")
		promptOptions(bill)
	}

}

func createBill() Bill {
	name, _ := getInput("Enter Customer Name: ")

	bill := newBill(name)
	fmt.Println("Created the bill - ", name)

	return bill
}

func main() {
	myBill := createBill()
	promptOptions(myBill)
	fmt.Println(myBill.format())
}
