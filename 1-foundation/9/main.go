package main

import "fmt"

func main() {
	fmt.Println(sum(1, 2, 42, 54, 2, 3, 52, 23))
}

func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}
