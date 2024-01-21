package main

import "fmt"

type person struct {
	fname string
	lname string
	// age   int
}

func main() {
	var p1 person
	p1.fname = "Ajith"
	p1.lname = "Reji"
	p1.printPerson()


}

func (p person) printPerson(){
	fmt.Printf("The person is %v %v", p.fname, p.lname)
}