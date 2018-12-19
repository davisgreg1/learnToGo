package main

import "fmt"

func main() {
	ids := []int{33, 76, 54, 23, 11, 2}

	// Loop throug ids
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// not using index
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	// Add all the ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum:", sum)

	//Range with map
	emails := map[string]string{"Greg": "Greg@gmail.com", "Carmelo": "carmelo@gmail.com"}
	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v)
	}
	for k := range emails {
		fmt.Println("Name: ", k)
	}
}
