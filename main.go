package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
}

func main() {
	p1 := person{"Arun", "Shekar", 30}
	fmt.Printf("Name: %v %v\n", p1.firstName, p1.lastName)
	fmt.Println(p1)

	p2 := person{firstName: "P2First", lastName: "P2Last", age: 35}
	fmt.Printf("Person2: %v %v\n", p2.firstName, p2.lastName)

	var p3 person // Go assigns a zero-value to fields in this struct
	fmt.Println(p3)
}
