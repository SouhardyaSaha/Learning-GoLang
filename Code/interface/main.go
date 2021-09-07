/*
1. Go interfaces are implicit.
2. 2 types =>
	i. Concrete : Struct, Map
	ii. Interface

*/

package main

import "fmt"

type bot interface {
	getGreeting() string
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

type englishBot struct{}
type spanishBot struct{}

func (englishBot) getGreeting() string {
	return "Hello there!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	eb.getGreeting()
	printGreeting(sb)
	printGreeting(eb)
}
