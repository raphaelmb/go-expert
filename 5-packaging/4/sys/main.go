package main

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/raphaelmb/go-expert/5-packaging/3/math"
)

func main() {
	m := math.Math{1, 2}
	fmt.Println(m.Add())
	fmt.Println(uuid.New().String())
}
