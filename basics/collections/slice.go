package main

import "fmt"

func main() {
	slice := []int{1, 2, 3}
	slice = append(slice, 5, 42, 21)

	fmt.Println(slice)

	// take all elements starts from index 1
	slice2 := slice[1:]

	// take 2 elements from the start
	slice3 := slice[:2]

	// take index 1 until 2
	slice4 := slice[1:2]
	fmt.Println(slice2, slice3, slice4)
}
