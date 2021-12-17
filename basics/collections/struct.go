package collections

import "fmt"

type user struct {
	ID        int
	FirstName string
	LastName  string
	GPA       float32
}

func main() {
	type user struct {
		ID        int
		FirstName string
		LastName  string
	}

	var u user
	u.ID = 123
	u.FirstName = "Samuel"
	u.LastName = "Toh"
	fmt.Println(u)

	// must end with a comma when doing muilti line of struct
	u2 := user{
		ID:        124,
		FirstName: "Samle",
		LastName:  "Toh",
	}

	fmt.Println(u2)
}
