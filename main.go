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
	fmt.Printf("Name: %v %v %+v\n", p1.firstName, p1.lastName, p1.contactInfo)
	fmt.Println(p1)

	p2 := person{firstName: "P2First", lastName: "P2Last", age: 35}
	fmt.Printf("Person2: %v %v %v\n", p2.firstName, p2.lastName, p2.age)

	var p3 person // Go assigns a zero-value to fields in this struct
	fmt.Println(p3)
}
