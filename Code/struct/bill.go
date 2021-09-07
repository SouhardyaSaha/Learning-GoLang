package main

import (
	"fmt"
	"os"
)

type Bill struct {
	name  string
	items map[string]float64
	tip   float64
}

func newBill(name string) Bill {
	bill := Bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}

	return bill
}

// receiver functions of struct
func (b Bill) format() string {
	fs := "Bill Break Down: \n"
	total := 0.0

	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ... $%v \n", k+":", v)
		total += v
	}

	fs += fmt.Sprintf("%-25v ... $%0.2f \n", "total:", total)

	return fs
}

func (bill *Bill) updateTip(tip float64) {

	// (*bill).tip = tip  // No need to deference struct is automatically dereferenced
	bill.tip = tip
}

func (bill *Bill) addItem(name string, price float64) {
	bill.items[name] = price
}

func (bill *Bill) save() {
	data := []byte(bill.format())
	err := os.WriteFile("bills/"+bill.name+".txt", data, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("Bill saved to file")
}
