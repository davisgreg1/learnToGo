package main

import "fmt"

func main() {
	// //Define map
	// emails := make(map[string]string)

	// // ASSIGN KEY VALS
	// emails["Greg"] = "greg@gmail.com"
	// emails["Sky"] = "sky@gmail.com"
	// emails["Semira"] = "Semira@gmail.com"
	// emails["Carmelo"] = "Carmelo@gmail.com"

	// Declare map and add key vals
	emails := map[string]string{"Greg": "Greg@gmail.com", "Carmelo": "carmelo@gmail.com"}
	fmt.Println(len(emails))

	//Delete from a map
	delete(emails, "Greg")
	fmt.Println(emails)
}
