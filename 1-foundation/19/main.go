package main

func main() {

	for i := 0; i < 10; i++ {
		println(i)
	}

	i := 0
	for i < 10 {
		println(i)
		i++
	}

	nums := []int{1, 2, 3, 4, 5}
	for i, v := range nums {
		println(i, v)
	}

	for {
		println("infinite")
	}
}
