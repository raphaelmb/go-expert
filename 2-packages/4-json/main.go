package main

import (
	"encoding/json"
	"os"
)

type Account struct {
	Number  int `json:"n"`
	Balance int `json:"b"`
}

func main() {
	account := Account{1, 100}
	res, err := json.Marshal(account)
	if err != nil {
		panic(err)
	}
	println(string(res))

	err = json.NewEncoder(os.Stdout).Encode(account)
	if err != nil {
		panic(err)
	}

	// pureJson := []byte(`{"Number": 2, "Balance": 200}`)
	pureJson := []byte(`{"n": 2, "b": 200}`)

	var accountX Account
	err = json.Unmarshal(pureJson, &accountX)
	if err != nil {
		panic(err)
	}

	println(accountX.Balance)
}
