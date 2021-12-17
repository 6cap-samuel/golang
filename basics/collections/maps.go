package collections

import "fmt"

func main() {

	// create a map with the KV of foo that maps to 42
	m := map[string]int{"foo": 42}
	fmt.Println(m)
	fmt.Println(m["foo"])

	// setting 27 into foo
	m["foo"] = 27
	fmt.Println(m)

	// remove the entry
	delete(m, "foo")
	fmt.Println(m)
}
