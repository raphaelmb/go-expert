package main

type Client struct {
	Name string
}

func (c *Client) Walk() {
	c.Name = "John Rambo"
	// fmt.Printf("The client %v walked\n", c.Name)
}

type Account struct {
	Balance int
}

func NewAccount() *Account {
	return &Account{0}
}

func (a *Account) Simulate(value int) int {
	a.Balance += value
	return a.Balance
}

func main() {
	client := Client{"John"}
	client.Walk()
	// fmt.Printf("The client's name is %v\n", client.Name)

	account := Account{100}
	account.Simulate(200)
	println(account.Balance)
}
