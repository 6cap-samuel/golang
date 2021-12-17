package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	fmt.Println("Printing from collections")
	slice := []int{1, 2, 3}
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	for index, value := range slice {
		println(index, value)
	}

	fmt.Println("")
	fmt.Println("Printing from map")
	wellKnownPorts := map[string]int{"http": 80, "https": 443}
	for key, value := range wellKnownPorts {
		println(key, value)
	}
}
