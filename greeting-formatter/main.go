package main

import (
	"fmt"
)

func Formatter(name, city string) string {
	greeting := fmt.Sprintf("Hello, %v from %v!", name, city)
	return greeting
}

func main() {
	greeting := Formatter("John Doe", "San Francisco")
	fmt.Println(greeting)
}
