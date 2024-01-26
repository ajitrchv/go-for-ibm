package main

import "fmt"

type bot interface {
	getGreeting() string
}

// type animal interface {
// 	getDetails(int, string) string
// }

// type human struct {
// 	name string
// 	age  int
// }

// type dog struct {
// 	name string
// 	age  int
// }

type engBot struct {
}

type espBot struct {
}

func main() {
	eb := engBot{}
	sb := espBot{}
	printGreeting(eb)
	printGreeting(sb)

	// elon := human{name: "musk", age: 55}

	// polo := dog{name: "marco", age: 1}

}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (eb engBot) getGreeting() string {
	// very custom logic for generating English greeting
	return "Hi there!"
}

func (sb espBot) getGreeting() string {
	// very custom logic for generating español greeting
	return "Hola!"
}
