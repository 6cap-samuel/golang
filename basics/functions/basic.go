package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Hello World")

	port := 3000
	_, err := startWebServer(port, 2)
	fmt.Println(err, "is true with code of ")
}

func startWebServer(
	port,
	numberOfRetires int,
) (int, error) {
	fmt.Println("Starting server on ", port)
	fmt.Println("Server started")

	return 2, errors.New("SOmething went wrong")

}
