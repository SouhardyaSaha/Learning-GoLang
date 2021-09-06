package main

import "fmt"

type Bill struct {
	name  string
	items map[string]float64
	tip   float64
}

func newBill(name string) Bill {
	bill := Bill{
		name:  name,
		items: map[string]float64{"cake": 1.34, "coke": 3.21},
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
