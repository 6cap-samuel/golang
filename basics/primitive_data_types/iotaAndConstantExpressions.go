package main

import (
	"fmt"
)

// iota starts with 0 and it increments for every constants
const (
	first = iota
	second
)

// iota resets in diff constant blocks
const (
	third = iota
	forth
)

func main() {
	fmt.Println(first, second, third, forth)
}
