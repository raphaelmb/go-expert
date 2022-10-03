package main

import "fmt"

func main() {
	salary := map[string]int{"T1": 1000, "T2": 2000, "T3": 3000}
	delete(salary, "T1")

	// sal := make(map[string]int)
	// sal2 := map[string]int{}

	for name, sal := range salary {
		fmt.Printf("%s's salary is %d\n", name, sal)
	}

	for _, sal := range salary {
		fmt.Printf("The salary is %d\n", sal)
	}
}
