package main

import (
	"fmt"
)

func main() {

	// strings
	var nameOne string = "mario"
	var nameTwo = "apple"
	var namethree string

	fmt.Println(
		nameOne,
		nameTwo,
		namethree,
	)

	nameFour := "newProcessOfStringAssinging"
	fmt.Println(nameFour)

	// Ints
	var ageOne int = 20
	var ageTwo = 30
	ageThree := 40

	fmt.Println(ageOne, ageTwo, ageThree)

	// Bits and memories
	var range1 int8 = 127
	var range2 int32 = 255

	fmt.Println(range1, range2)

	// format specifiers strings with var in one line
	fmt.Printf("my age is %v and my name is %v \n", ageOne, nameOne)

	// saving formatted string into var
	savedString := fmt.Sprintf("my age is %v and my name is %v", ageOne, nameOne)
	fmt.Println(savedString)

}
