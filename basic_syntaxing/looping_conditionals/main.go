package main

import (
	"fmt"
)

func main() {

	names := []string{"a", "BN", "q23"}

	// for i := 0; i < len(names); i++ {
	// 	fmt.Println(names[i])
	// }

	for index, value := range names {
		if index == 1 {
			fmt.Println("At One")
			continue
		} else {
			fmt.Printf("Value of %v is %v \n", index, value)
		}

	}

}
