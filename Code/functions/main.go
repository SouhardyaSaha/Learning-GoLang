package main

import (
	"fmt"
	"math"
)

func sayGreetings(n string) {
	fmt.Printf("Hello %v \n", n)
}

func sayBye(n string) {
	fmt.Printf("Bye %v \n", n)
}

func cycleNames(names []string, f func(string)) {
	for _, v := range names {
		f(v)
	}
}

func circleArea(r float64) float64 {
	return math.Pi * r * r
}

func main() {
	names := []string{"sa", "asdf123"}

	cycleNames(names, sayGreetings)
	cycleNames(names, sayBye)

	fmt.Printf("The area of circle is %0.3f", circleArea(10.5))
}
