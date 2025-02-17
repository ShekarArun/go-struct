package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}
type person struct {
	firstName string
	lastName  string
	age       int
	contactInfo
}

func main() {
	p1 := person{"Arun", "Shekar", 30, contactInfo{"a@a.com", 34}}
	p1.print()
	p1.updateFirstName("NewFirst")
	p1.print()

	p2 := person{firstName: "P2First", lastName: "P2Last", age: 35}
	p2.print()

	var p3 person // Go assigns a zero-value to fields in this struct
	p3.print()
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func (p person) updateFirstName(n string) {
	p.firstName = n
	fmt.Printf("First name updated to: %v\n", p.firstName)
}
