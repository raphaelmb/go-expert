package main

import "github.com/raphaelmb/go-expert/7-api/configs"

func main() {
	config, _ := configs.LoadConfig(".")
	println(config.DBDriver)
}
