package main

import "fmt"

func main() {
	// MAIN TYPES
	// string
	// bool
	// int
	// int int8 int16 int32 int64
	// uint uint8 uint16 uint32 uint64 (0...INFINITY no negative)
	// byte - alias for uint8
	// rune - alias for int32
	// float32 float64(this is the one we usually use)
	// complex64 complex128

	// Using var

	// var name = "Greg"
	var age int32 = 33
	// var name = "Greg"
	const isCool = true
	var size float32 = 2.3

	//Shorthand way to make a variable
	// name := "Greg"
	// email := "davisgreg1@gmail.com"

	name, email := "Greg", "davisgreg1@gmail.com"
	// size := 1.3
	fmt.Println(name, age, isCool, email)
	fmt.Printf("%T\n", size)
}
