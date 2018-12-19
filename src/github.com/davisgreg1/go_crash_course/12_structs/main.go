package main

import (
	"fmt"
	"strconv"
)

// Person is Define Person struct
type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

// Greeting method (value receiver because we arent going to change anything; just use the values and return them)
func (p Person) greet() string { //the 'p' is similar to the 'this' keyword
	return "Hello, my name is:" + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

// hasBirthday method (pointer receiver because we are going to change something. Because its a pointer receiver we must use the asteriks)
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried (pointer receiver)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "M" {
		return
	} else {
		p.lastName = spouseLastName
	}
}

func main() {
	// Init person using struct
	person1 := Person{firstName: "Greg", lastName: "Davis", city: "NYC", gender: "M", age: 33}

	// Alternative
	person2 := Person{"Ty", "Yonas", "Africa", "F", 33}

	// fmt.Println(person1)

	// //Get single field
	// fmt.Println(person1.firstName)
	// fmt.Println(person1.age)
	// person1.age++
	// fmt.Println(person1.age)

	person1.hasBirthday()
	person2.getMarried("Davis")
	fmt.Println(person1.greet())
	fmt.Println(person2.greet())
}
