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

func circleAreaAndCircumference(r float64) (float64, float64) {
	return math.Pi * r * r, 2 * math.Pi * r
}

func main() {
	// names := []string{"sa", "asdf123"}

	// cycleNames(names, sayGreetings)
	// cycleNames(names, sayBye)

	// fmt.Printf("The area of circle is %0.3f", circleArea(10.5))

	// return multiple values in a function
	area, circumference := circleAreaAndCircumference(10.5)
	fmt.Printf("The area of circle is %0.3f and Circumference is %0.3f", area, circumference)

}
