package main

import (
	"fmt"
)

func main() {
	var firstName *string = new(string)
	*firstName = "Samuel"
	fmt.Println(*firstName)

	var secondName string = "TOH"
	fmt.Println(&secondName, secondName)

	var thirdName string = "Micheal"
	var changing = &thirdName
	*changing = "Silly"
	fmt.Println(*changing)
	fmt.Println(thirdName)
}
