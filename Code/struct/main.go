package main

import (
	"fmt"
)

func main() {
	myBill := newBill("Mario's Bill")
	fmt.Println(myBill.format())
	// myBill.format()

}
