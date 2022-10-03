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
	fmt.Printf("e's type is %T\n", e)
	fmt.Printf("e's value is %v\n", e)
	fmt.Println()
	fmt.Printf("f's type is %T\n", f)
}
