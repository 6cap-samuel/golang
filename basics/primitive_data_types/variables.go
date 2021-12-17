package main

import (
	"fmt"
)

func main() {
	// all variables that are being declared must be used else it will face compilation error.
	// explicit init - with type declarations
	var i int
	i = 42
	fmt.Println(i)

	var f float32 = 3.14
	fmt.Println(f)

	// implicit init - without type declarations
	firstName := "Samuel"
	fmt.Println(firstName)

	b := true
	fmt.Println(b)

	c := complex(3, 4)
	fmt.Println(c)

	r, im := real(c), imag(c)
	fmt.Println(r)
	fmt.Println(im)
}
