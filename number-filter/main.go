package main

import "fmt"

func Filter(numbers []int) []int {
	evenNumbers := make([]int, 0)
	for _, v := range numbers {
		if v%2 == 0 {
			evenNumbers = append(evenNumbers, v)
		}
	}
	return evenNumbers
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	evenNumbers := Filter(numbers)
	printStatement := fmt.Sprintf("Original: %v\nFiltered: %v", numbers, evenNumbers)
	fmt.Println(printStatement)
}
