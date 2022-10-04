package main

import "fmt"

type Client struct {
	Name   string
	Age    int
	Active bool
}

func main() {
	client := Client{
		Name:   "John",
		Age:    30,
		Active: true,
	}

	client.Active = false

	fmt.Printf("Name: %s, age: %d, active: %t\n", client.Name, client.Age, client.Active)

}
