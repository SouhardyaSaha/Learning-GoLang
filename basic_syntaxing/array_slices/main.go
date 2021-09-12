package main

import (
	"fmt"
)

func main() {

	// var ages [3]int = [3]int{1, 2, 3}
	var ages = [3]int{1, 2, 3}

	fmt.Println(ages, len(ages))

	// slices (uses arrays under the hood)
	// length of slices can be changed but arrays can't be changed.
	var names = []string{"Dip", "Bala", "Rumi"}
	fmt.Println(names)

	// using range in slice
	fmt.Println(names[1:3]) // 1 to 2
	fmt.Println(names[:3])  // first to 2
	fmt.Println(names[1:])  // 1 to last

	// appned in slice
	names = append(names, "Ruby")
	fmt.Println(names)

}
