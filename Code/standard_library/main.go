package main

import (
	"fmt"
	"sort"
)

func main() {

	// greeting := "Hello from the other side"

	// fmt.Println(strings.Count(greeting, "h"))
	// fmt.Println(strings.Contains(greeting, "Hello"))
	// fmt.Println(strings.Split(greeting, " "))

	ages := []int{12, 13, 56, 63, 6, 72}
	sort.Ints(ages)
	fmt.Println(ages)
}
