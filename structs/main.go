package main

import "fmt"

type person struct {
	fname string
	lname string
	// age   int
}

func main() {
	person1 := person{fname: "Ajith", lname: "Reji"}
	fmt.Println(person1)
}