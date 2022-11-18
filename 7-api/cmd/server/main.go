package main

import "github.com/raphaelmb/go-expert/7-api/configs"

func main() {
	_, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
}
