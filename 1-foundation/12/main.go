package main

import "fmt"

type Address struct {
	Street string
	Number int
	City   string
	State  string
}

type Client struct {
	Name   string
	Age    int
	Active bool
	Address
}

func (c *Client) Deactivate() {
	c.Active = false
	fmt.Printf("The client %s was deactivated\n", c.Name)
}

func main() {
	client := Client{
		Name:   "John",
		Age:    30,
		Active: true,
	}

	client.Deactivate()

}
