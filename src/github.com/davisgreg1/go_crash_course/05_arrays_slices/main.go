package main

import "fmt"

func main() {
	// Arrays
	// var fruitArr [2]string
	// Assign Vals
	// fruitArr[0] = "Apple"
	// fruitArr[1] = "Banana"

	// Declare and Assign shortHand
	// fruitArr := [2]string{"Apple", "Banana"}

	fruitSlice := []string{"Apple", "Orange", "Banana"}

	fmt.Println(fruitSlice)
	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[1:2])
}
