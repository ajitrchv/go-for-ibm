package main

import "fmt"

type bot interface {
	getGreeting() string
}

type engBot struct {
}

type espBot struct {
}

type wannaBot struct {
	count int
}

func main() {
	eb := engBot{}
	sb := espBot{}
	wb := wannaBot{}

	printGreeting(eb)
	printGreeting(sb)
	printGreeting(wb)

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

func (wb wannaBot) getGreeting() string {
	return "Here is this wanna bot!"
}
