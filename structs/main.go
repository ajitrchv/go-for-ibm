package main

import "fmt"

type contactInfo struct{
	email string
	zipCode int
}

type person struct {
	fname string
	lname string
	contact contactInfo
	// age   int
}

func main() {
	var p1 person
	p1.fname = "Ajith"
	p1.lname = "Reji"
	p1.contact.email="ajithvreji@github.com"
	p1.contact.zipCode=685605
	p1.printPerson()


	hardpope := person{
		fname: "hard",
		lname:  "pope",
		contact: contactInfo{
			email: "hardpope@hp.com",
			zipCode: 000011,
		},
	}

	hardpope.printContact()

	h := &hardpope

	h.updateName("hardy")

	hardpope.printPerson()

}

func (p person) printPerson(){
	fmt.Printf("The person is %v %v\n", p.fname, p.lname)
}

func (p person) printContact(){
	fmt.Printf("The contact of %v is %+v\n", p.fname, p.contact)
}

func (ptr *person) updateName(fname string){
	ptr.fname = fname
	fmt.Println("Name updated!")
}