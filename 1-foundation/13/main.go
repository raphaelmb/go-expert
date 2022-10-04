package main

import "fmt"

type Address struct {
	Street string
	Number int
	City   string
	State  string
}

type Entity interface {
	Deactivate()
}

type Company struct {
	Name string
}

func (c Company) Deactivate() {}

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

func Deactivation(e Entity) {
	e.Deactivate()
}

func main() {
	client := Client{
		Name:   "John",
		Age:    30,
		Active: true,
	}

	company := Company{"Company"}

	Deactivation(company)

	Deactivation(&client)

}
