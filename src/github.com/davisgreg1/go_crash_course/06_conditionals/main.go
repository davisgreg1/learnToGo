package main

import "fmt"

func main() {
	x := 10
	y := 10
	// if else
	if x <= y {
		fmt.Printf("%d is less than or equal to %d\n", x, y)
	} else {
		fmt.Printf("%d is less than %d \n", y, x)
	}

	//else if
	color := "red"

	if color == "red" {
		fmt.Println("color is red")
	} else if color == "blue" {
		fmt.Println("color is blue")
	} else {
		fmt.Println("color is not blue or red")
	}

	// Switch

	switch color {
	case "red":
		fmt.Println("color is red")
	case "blue":
		fmt.Println("color is blue")
	default:
		fmt.Println("color is not blue or damn red")
	}
}
