package main

import "fmt"

const a = "Hello, World!"

type ID int

var (
	b bool    = true
	c int     = 10
	d string  = "Go"
	e float64 = 1.2
	f ID      = 1
)

func main() {
	var array [3]int
	array[0] = 10
	array[1] = 20
	array[2] = 30

	for i, v := range array {
		fmt.Printf("value of index %d is %d\n", i, v)
	}
}
