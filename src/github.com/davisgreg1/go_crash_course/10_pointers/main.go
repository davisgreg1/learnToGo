package main

import "fmt"

func main() {
	// a pointer alloww you to point to the memory address or location of a value thats in a variable

	a := 5
	b := &a

	fmt.Println(a, b)
	fmt.Printf("%T\n", a) //int
	fmt.Printf("%T\n", b) //a *int represents a pointer

	// use * to read value from an address
	fmt.Println(*b)
	fmt.Println(*&a)

	// Change val with pointer
	*b = 10
	fmt.Println("a is:", a)
	fmt.Println("b is:", b)
}
