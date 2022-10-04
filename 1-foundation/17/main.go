package main

type MyNumber int

type Number interface {
	~int | float64
}

func Sum[T Number](m map[string]T) T {
	var total T
	for _, v := range m {
		total += v
	}
	return total
}

func Compare[T comparable](a T, b T) bool {
	if a == b {
		return true
	}
	return false
}

func main() {
	m := map[string]int{"John": 1, "Jack": 2, "Clint": 3}
	m2 := map[string]float64{"John": 1.1, "Jack": 2.2, "Clint": 3.3}
	m3 := map[string]MyNumber{"John": 1, "Jack": 2, "Clint": 3}
	println(Sum(m))
	println(Sum(m2))
	println(Sum(m3))

	println(Compare(10, 10))
}
