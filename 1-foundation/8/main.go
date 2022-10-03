package main

import (
	"errors"
	"fmt"
)

func main() {
	value, err := sum(5, 1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("The value is", value)
}

func sum(a, b int) (int, error) {
	if a+b >= 50 {
		return 0, errors.New("Sum is greater than 50")
	}
	return a + b, nil
}
