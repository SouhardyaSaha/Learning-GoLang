package main

import "fmt"

func updateNameByRef(name *string) {
	*name = "Virat"
}

func main() {
	name := "Sachin"
	memory_address := &name

	// fmt.Println(memory_address)
	// fmt.Println("Value of memory address: ", *memory_address)

	fmt.Println(name)
	updateNameByRef(memory_address)
	fmt.Println(name)

}
