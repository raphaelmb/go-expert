package main

import "fmt"

func main() {
	var variable interface{} = "John Rambo"
	println(variable.(string))
	res, ok := variable.(int)
	fmt.Printf("Res value is %v and ok value is %v\n", res, ok)
}
